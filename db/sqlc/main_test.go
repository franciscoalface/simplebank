package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/franciscoalface/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load env configs:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}