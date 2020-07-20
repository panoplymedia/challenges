package postgres_test

import (
	"flag"
	"os"
	"testing"
)

var connectionArgs = "user=postgres password=test host=postgres port=5432 database=postgres sslmode=disable"

func TestMain(m *testing.M) {
	flag.StringVar(&connectionArgs,
		"db-configs",
		"user=postgres password=test host=postgres port=5432 database=postgres sslmode=disable",
		"the configuration string for connecting to the database")
	flag.Parse()
	os.Exit(m.Run())
}
