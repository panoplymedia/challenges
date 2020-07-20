package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"

	"github.com/panoplymedia/sales_admin/cmd/server/internal/service"
	"github.com/panoplymedia/sales_admin/internal/postgres"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		runErr, ok := err.(cmdError)
		if !ok {
			os.Exit(exitUnknown)
		}

		os.Exit(runErr.code)
	}
}

func run(args []string, stdOut io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	listenAddress := flags.String("listen-address", ":8888", "the port on which the server should listen")
	dbConfigs := flags.String("db-configs", "", "the configuration string for connecting to the database")
	shutdownTimeout := flags.Duration("shutdown-timeout", time.Second*30, "the duration to wait when shutting down the server")
	templatesGlob := flags.String("templates", "/srv/views/*.html", "the file glob to use when searching for templates")
	staticFilesDir := flags.String("static-files", "/srv/assets", "the path to the directory that contains the static asset files")
	logLevel := logLevel(log.ERROR)
	flags.Var(&logLevel, "log-level", "log level to use: DEBUG, INFO, WARN, ERROR, or OFF")

	err := flags.Parse(args[1:])
	if err != nil {
		return cmdError{code: badFlags, err: err}
	}

	db, err := postgres.Connect(*dbConfigs)
	if err != nil {
		return cmdError{code: failedSetup, err: errors.Wrap(err, "failed to connect to the database")}
	}
	defer db.Close()

	logger := log.New("server")
	logger.SetLevel(log.Lvl(logLevel))

	errs := make(chan error, 1)
	defer close(errs)
	// Start the service.
	service, err := service.New(
		service.WithSalesService(&postgres.SalesService{DB: db}),
		service.WithTemplates(*templatesGlob),
		service.WithStaticFiles(*staticFilesDir),
		service.WithLogger(logger),
	)
	if err != nil {
		return cmdError{code: failedServerSetup, err: errors.Wrap(err, "failed to setup the server")}
	}
	go func() {
		errs <- service.Start(*listenAddress)
	}()

	// Listen for SIGTERM so that we can shutdown cleanly.
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGKILL)

	// Wait here until we get an error or a SIGTERM.
	select {
	case err := <-errs:
		return cmdError{code: serverError, err: errors.Wrap(err, "the service failed")}
	case <-signals:
		ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
		defer cancel()

		err := service.Shutdown(ctx)

		// Wait for the service to write channel so that it can be cleaned up.
		<-errs
		if err != nil {
			return cmdError{code: shutdownError, err: errors.Wrap(err, "error occurred while shutting down")}
		}
	}

	return nil
}
