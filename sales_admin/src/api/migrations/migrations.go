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
	workDir, _ := os.Getwd()
	s, err := f.Open("file://" + path.Join(workDir, migrationsDir))
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

	if err := createAdminUser(db, log); err != nil {
		log.Fatalf("Migrations: %+v", err)
	}
}

func createAdminUser(db *sql.DB, log *logrus.Logger) error {
	email := os.Getenv("ADMIN_USER")
	pass := os.Getenv("ADMIN_PASS")
	_, err := db.Exec(`insert into "user" (password, email, role)
		values (crypt($1, gen_salt('bf', 8)), $2, 'admin') on conflict do nothing;`, pass, email)

	return err
}
