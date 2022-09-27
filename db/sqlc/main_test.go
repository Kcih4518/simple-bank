package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// Queries contain a `dbtx` : a db connection or a transaction
var testQueries *Queries

// the main entry point off all unit tests
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// New is defined in the db.go
	testQueries = New(conn)

	// m.Run(): star running unit-tests
	os.Exit(m.Run())
}