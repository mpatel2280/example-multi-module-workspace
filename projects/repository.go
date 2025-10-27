package projects

import (
	"database/sql"
	"log"

	"example-multi-module-workspace/db"
)

func GetAll() ([]Project, error) {
	var projects []Project

	query := "SELECT id, name, tech_stack FROM projects"

	err := db.Select(query, nil, func(rows *sql.Rows) error {
		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.TechStack); err != nil {
			return err
		}
		projects = append(projects, p)
		return nil
	})

	if err != nil {
		log.Printf("GetAll() error: %v", err)
		return nil, err
	}

	return projects, nil
}
