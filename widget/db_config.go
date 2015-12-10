package widget

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/dakiva/dbx"
	"github.com/jmoiron/sqlx"
)

const (
	POSTGRES_TYPE = "postgres"
)

// these database objects should not be statically referenced by any repositories, rather they should be passed in during construction of a repository.
var db *sqlx.DB
var queryMap dbx.QueryMap

type DBConfig struct {
	Hostname           string `json:"hostname"`
	Port               int    `json:"port"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
	DBName             string `json:"dbname"`
	SSLMode            string `json:"sslmode"`
	ConnectTimeout     string `json:"connectTimeout"`
	SchemaName         string `json:"schema_name"`
	RolePassword       string `json:"role_password"`
	QueryFile          string `json:"query_file"`
}

func (this *DBConfig) toDSN() string {
	dsn := ""
	if this.Hostname != "" {
		dsn += "host=" + this.Hostname + " "
	}
	if this.Port > 0 {
		dsn += fmt.Sprintf("port=%d ", this.Port)
	}
	dsn += "user=" + this.SchemaName + " "
	if this.RolePassword != "" {
		dsn += "password=" + this.RolePassword + " "
	}
	dsn += "dbname=" + this.DBName + " "
	dsn += "sslmode=" + this.SSLMode + " "
	if this.ConnectTimeout != "" {
		dsn += "connect-timeout=" + this.ConnectTimeout
	}
	return strings.TrimSpace(dsn)
}

func (this *DBConfig) Validate() error {
	if this.DBName == "" {
		return errors.New("Database name must be specified.")
	}
	if this.SSLMode == "" {
		return errors.New("Database SSLMode must be specified.")
	}
	if this.SchemaName == "" {
		return errors.New("Schema name must be specified.")
	}
	if this.RolePassword == "" {
		return errors.New("Schema role password must be specified.")
	}
	if this.QueryFile == "" {
		return errors.New("A valid query file path must be specified.")
	}
	return nil
}

func (this *DBConfig) Initialize(schemaVersion int) error {
	err := InitDB(this.toDSN(), this.SchemaName, schemaVersion, this.MaxOpenConnections, this.MaxIdleConnections)
	if err != nil {
		return err
	}

	err = InitQueryMap(this.QueryFile)
	if err != nil {
		return err
	}
	return nil
}

// initialize and cache the database instance
func InitDB(dsn, schemaName string, expectedSchemaVersion, maxOpenConnections, maxIdleConnections int) error {
	var err error
	db, err = initDB(dsn, schemaName, expectedSchemaVersion, maxOpenConnections, maxIdleConnections)
	if err != nil {
		return err
	}
	return nil
}

// initialize and cache all named SQL queries
func InitQueryMap(queryFiles ...string) error {
	var err error
	queryMap, err = dbx.LoadNamedQueries(queryFiles...)
	if err != nil {
		return err
	}
	return nil
}

// returns the DB instance
func GetDB() *sqlx.DB {
	return db
}

// returns the map of all named queries
func GetQueryMap() dbx.QueryMap {
	return queryMap
}

func initDB(dsn, schemaName string, expectedSchemaVersion, maxOpenConnections, maxIdleConnections int) (*sqlx.DB, error) {
	db, err := sqlx.Connect(POSTGRES_TYPE, dsn)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	if err != nil {
		return nil, err
	}
	version, err := dbx.GetCurrentSchemaVersion(schemaName, db)
	if err != nil {
		return nil, err
	}
	if int64(expectedSchemaVersion) != version {
		return nil, errors.New(fmt.Sprintf("Schema version mismatch: %v != %v", version, expectedSchemaVersion))
	}
	return db, nil
}

// grab the sequential id from the result rows returned from an INSERT...RETURNING query. This function closes the result set.
func GetNextId(rows *sqlx.Rows) (int64, error) {
	defer rows.Close()
	if rows.Next() {
		var id sql.NullInt64
		err := rows.Scan(&id)
		if err != nil {
			return -1, err
		}
		if id.Valid {
			return id.Int64, nil
		}
	}
	return -1, errors.New("No rows inserted or returned.")
}
