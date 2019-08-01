package migrations

import (
	"database/sql"
	"os"
	"path"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source"
	"github.com/mattes/migrate/source/file"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func loadMigrations(migrationsDir string, log *logrus.Logger) source.Driver {
	f := &file.File{}
	s, err := f.Open("file://" + path.Join(os.Getenv("PROJECT_ROOT"), migrationsDir))
	if err != nil {
		log.Fatalf("Error loading migrations %+v", err)
	}
	return s
}

func Migrate(db *sql.DB, log *logrus.Logger) {
	src := loadMigrations("migrations", log)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating driver %+v", err)
	}

	m, err := migrate.NewWithInstance("go-bindata", src, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migrations: %+v", err)
	}
}
