package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Kcih4518/simple-bank/util"
	_ "github.com/lib/pq"
)

// Queries contain a `dbtx` : a db connection or a transaction
var testQueries *Queries
var testDB *sql.DB

// the main entry point off all unit tests
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// New is defined in the db.go
	testQueries = New(testDB)

	// m.Run(): star running unit-tests
	os.Exit(m.Run())
}
