package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	// "github.com/lib/pq" must be imported for the postgres driver
	_ "github.com/lib/pq"
)

var connectionPool *sqlx.DB

// CreateConnectionPool created a new database connection pool
func CreateConnectionPool() error {

	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Println(dbInfo)

	db, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(50)

	connectionPool = db

	return nil
}

// NewConnection returns a db connection
func NewConnection() (*sqlx.DB, error) {

	if connectionPool == nil {
		err := CreateConnectionPool()
		if err != nil {
			return nil, err
		}
	}

	return connectionPool, nil
}

// CreateSaleTable creates the sale database table
func CreateSaleTable() error {

	conn, err := NewConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS sale (
    custome_name TEXT NOT NULL,
    item_description TEXT NOT NULL,
    item_price DOUBLE PRECISION NOT NULL,
    quantity INTEGER NOT NULL,
    merchant_name TEXT NOT NULL,
    merchant_address TEXT NOT NULL
		);`)
	if err != nil {
		return err
	}

	return err
}
