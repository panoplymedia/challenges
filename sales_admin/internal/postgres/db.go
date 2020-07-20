package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Connect connects to a Postgres database using the connection string.
func Connect(args string) (*sql.DB, error) {
	return sql.Open("pgx", args)
}
