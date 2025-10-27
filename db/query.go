package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Select runs a SELECT query and applies a row handler callback
// It is generic so any table/entity can reuse it.
func Select(query string, args []any, scanFunc func(*sql.Rows) error) error {
	if dbInstance == nil {
		return fmt.Errorf("database not connected — call db.Connect() first")
	}

	rows, err := dbInstance.Query(query, args...)
	if err != nil {
		return fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := scanFunc(rows); err != nil {
			log.Printf("row scan error: %v", err)
		}
	}

	return rows.Err()
}
