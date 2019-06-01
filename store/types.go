package store

import (
	"database/sql"
)

// Database contains a SQL database
type Database struct {
	Instance *sql.DB
}
