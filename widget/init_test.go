package widget

import (
	"fmt"
	"os"
	"testing"

	"github.com/dakiva/dbx"
	"github.com/jmoiron/sqlx"
)

var testdb *sqlx.DB
var testQueryMap dbx.QueryMap

// initializes repository tests by loading the DB and all named queries, panicing on error
func TestMain(m *testing.M) {
	schema := dbx.GenerateSchemaName("homdna")
	testdb = dbx.MustInitializeTestDB(schema, "../db/migrations")
	testQueryMap = dbx.MustLoadNamedQueries("../db/queries/widget-queries.json")
	checkSchemaVersion(schema, testdb)

	returnCode := m.Run()

	dbx.TearDownTestDB(schema)
	os.Exit(returnCode)
}

func checkSchemaVersion(schema string, db *sqlx.DB) {
	version, err := dbx.GetCurrentSchemaVersion(schema, db)
	if err != nil {
		panic("Could not retrieve schema version information from the database.")
	}
	if version != SCHEMA_VERSION {
		panic(fmt.Sprintf("The schema version %d != %d", version, SCHEMA_VERSION))
	}
}
