package service

import (
	"context"
	html_template "html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"

	"github.com/panoplymedia/sales_admin/internal/sales"
)

// config is used to configure the service. It's fields are set using Option
// functions.
type config struct {
	salesService   sales.Service
	logger         echo.Logger
	templates      *html_template.Template
	staticFilesDir string
}

// Option is used to configure the Service.
type Option func(*config) error

// WithSalesService configures the Service to use the given sales.Service.
func WithSalesService(s sales.Service) Option {
	return func(config *config) error {
		config.salesService = s

		return nil
	}
}

// WithLogger configures the Service to use the given log level.
func WithLogger(logger echo.Logger) Option {
	return func(config *config) error {
		config.logger = logger

		return nil
	}
}

// WithTemplates configures the Service to use the given templates found at the
// give path glob.
func WithTemplates(pattern string) Option {
	return func(c *config) error {
		templates, err := html_template.ParseGlob(pattern)
		if err != nil {
			return errors.Wrapf(err, "failed to parse the templates at %s", pattern)
		}
		c.templates = templates

		return nil
	}
}

// WithStaticFiles configures the service to service static files from the given
// directory.
func WithStaticFiles(directoryPath string) Option {
	return func(c *config) error {
		c.staticFilesDir = directoryPath

		return nil
	}
}

// Service is the http service that handles sales data.
type Service struct {
	sales sales.Service
	echo  *echo.Echo
}

// New Service returns a service for handling sales data.
func New(options ...Option) (*Service, error) {
	config := config{
		logger: log.New("sales_admin"),
	}
	for _, option := range options {
		option(&config)
	}

	switch {
	case config.salesService == nil:
		return nil, errors.New("sales service is not set, WithSalesService must be called")

	case config.templates == nil:
		return nil, errors.New("templates are not set, WithTemplates must be called")
	}

	e := echo.New()
	e.Logger = config.logger
	e.Renderer = (*template)(config.templates)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: config.logger.Output(),
	}))
	if config.staticFilesDir != "" {
		e.Static("/assets", config.staticFilesDir)
	}

	s := Service{
		sales: config.salesService,
		echo:  e,
	}
	s.setupRoutes()

	return &s, nil
}

// Start starts the http service and listens at the given address for incoming
// requests.
func (s *Service) Start(address string) error {
	return s.echo.Start(address)
}

// Shutdown stops the service gracefully or until the context timeouts or is canceled.
func (s *Service) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

// HTTPHandler returns the underlying http.Handler. This is useful for
// testing. Use service.Start to start the service.
func (s *Service) HTTPHandler() http.Handler {
	return s.echo
}
