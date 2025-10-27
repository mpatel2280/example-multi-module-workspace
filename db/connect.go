package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

// Connect returns a reusable MySQL *sql.DB instance
func Connect() *sql.DB {
	// Change connection details as needed
	dsn := "root:root@tcp(localhost:3306)/testdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL connect error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("MySQL ping failed: %v", err)
	}

	fmt.Println("✅ Connected to MySQL successfully")
	dbInstance = db
	return db
}

// Query executes a SQL query and returns the rows
func Query(query string, args ...any) (*sql.Rows, error) {
	if dbInstance == nil {
		return nil, fmt.Errorf("database not connected")
	}
	return dbInstance.Query(query, args...)
}
