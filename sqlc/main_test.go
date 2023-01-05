package sqlc

import (
	"database/sql"
	"github.com/earlofurl/pxthc/config"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	loadConfig, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load loadConfig:", err)
	}

	testDB, err = sql.Open(loadConfig.DBDriver, loadConfig.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
