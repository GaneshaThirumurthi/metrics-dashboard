package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GaneshaThirumurthi/metrics-dashboard/config"
	"github.com/GaneshaThirumurthi/metrics-dashboard/consts"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var server = config.DatabaseServer
var port = consts.DatabasePort
var user = config.DatabaseUsername
var password = config.DatabasePassword
var database = consts.DatabaseName

// StartServer starts DB server
func (db *Database) StartServer() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db.Instance, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.Instance.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	db.SelectVersion()
}

// SelectVersion gets and prints SQL Server version
func (db *Database) SelectVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.Instance.PingContext(ctx)
	if err != nil {
		log.Fatal("error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.Instance.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
