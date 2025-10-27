package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// findMigrationsDir finds the migrations directory by checking multiple possible locations
func findMigrationsDir() (string, error) {
	cwd, _ := os.Getwd()

	const migrationsDir = "migrations"

	possiblePaths := []string{
		"db/" + migrationsDir,
		"./db/" + migrationsDir,
		filepath.Join(cwd, "db", migrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), migrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), "..", migrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), "..", "..", "db", migrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), "..", "..", "..", "db", migrationsDir),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			log.Printf("Found migrations directory at: %s", path)
			return path, nil
		}
	}

	// Print debug information
	log.Printf("Current working directory: %s", cwd)
	log.Printf("Executable path: %s", os.Args[0])
	log.Printf("Executable directory: %s", filepath.Dir(os.Args[0]))
	return "", fmt.Errorf("migrations directory not found in any of the expected locations. Tried: %v", possiblePaths)
}

// findRefreshMigrationsDir finds the migrations directory by checking multiple possible locations
func findRefreshMigrationsDir() (string, error) {
	cwd, _ := os.Getwd()

	const refreshMigrationsDir = "refresh-migrations"

	possiblePaths := []string{
		"db/" + refreshMigrationsDir,
		"./db/" + refreshMigrationsDir,
		filepath.Join(cwd, "db", refreshMigrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), refreshMigrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), "..", refreshMigrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), "..", "..", "db", refreshMigrationsDir),
		filepath.Join(filepath.Dir(os.Args[0]), "..", "..", "..", "db", refreshMigrationsDir),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			log.Printf("Found migrations directory at: %s", path)
			return path, nil
		}
	}

	// Print debug information
	log.Printf("Current working directory: %s", cwd)
	log.Printf("Executable path: %s", os.Args[0])
	log.Printf("Executable directory: %s", filepath.Dir(os.Args[0]))
	return "", fmt.Errorf("migrations directory not found in any of the expected locations. Tried: %v", possiblePaths)
}

// RunMigrations executes all SQL migration files in the migrations directory
func RunMigrations() error {
	if dbInstance == nil {
		return fmt.Errorf("database not connected — call db.Connect() first")
	}

	// Find the migrations directory
	migrationsDir, err := findMigrationsDir()
	if err != nil {
		return err
	}

	// Read migration files
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	// Execute each migration file
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		filePath := filepath.Join(migrationsDir, file.Name())
		log.Printf("Running migration: %s", file.Name())

		// Read the SQL file
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file.Name(), err)
		}

		// Execute the SQL statements
		// Split by semicolon to handle multiple statements
		statements := strings.Split(string(content), ";")
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			_, err := dbInstance.Exec(stmt)
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", file.Name(), err)
			}
		}

		log.Printf("✅ Migration completed: %s", file.Name())
	}

	return nil
}

// RunMigrations executes all SQL migration files in the migrations directory
func RunRefreshMigrations() error {
	if dbInstance == nil {
		return fmt.Errorf("database not connected — call db.Connect() first")
	}

	// Find the migrations directory
	migrationsDir, err := findRefreshMigrationsDir()
	if err != nil {
		return err
	}

	// Read migration files
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	// Execute each migration file
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		filePath := filepath.Join(migrationsDir, file.Name())
		log.Printf("Running migration: %s", file.Name())

		// Read the SQL file
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file.Name(), err)
		}

		// Execute the SQL statements
		// Split by semicolon to handle multiple statements
		statements := strings.Split(string(content), ";")
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			_, err := dbInstance.Exec(stmt)
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", file.Name(), err)
			}
		}

		log.Printf("✅ Migration completed: %s", file.Name())
	}

	return nil
}

// SeedDatabase inserts sample data into the projects table
func SeedDatabase() error {
	if dbInstance == nil {
		return fmt.Errorf("database not connected — call db.Connect() first")
	}

	// Check if projects table already has data
	var count int
	err := dbInstance.QueryRow("SELECT COUNT(*) FROM projects").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing data: %v", err)
	}

	if count > 0 {
		log.Printf("Projects table already contains %d records. Skipping seed.", count)
		return nil
	}

	// Seed data
	seedProjects := []struct {
		name      string
		techStack string
	}{
		{"E-Commerce Platform", "Go, PostgreSQL, React"},
		{"Real-time Chat Application", "Go, WebSocket, Vue.js"},
		{"Data Analytics Dashboard", "Go, MySQL, Angular"},
		{"Mobile API Gateway", "Go, gRPC, Kubernetes"},
		{"Content Management System", "Go, MongoDB, Next.js"},
		{"IoT Device Manager", "Go, MQTT, Flutter"},
		{"Microservices Framework", "Go, Docker, Kubernetes"},
		{"Machine Learning Pipeline", "Go, Python, TensorFlow"},
		{"Payment Processing System", "Go, MySQL, React"},
		{"Social Media Backend", "Go, Redis, PostgreSQL"},
	}

	// Insert seed data
	for _, project := range seedProjects {
		_, err := dbInstance.Exec(
			"INSERT INTO projects (name, tech_stack) VALUES (?, ?)",
			project.name,
			project.techStack,
		)
		if err != nil {
			return fmt.Errorf("failed to insert seed data: %v", err)
		}
	}

	log.Printf("✅ Seeded %d projects into the database", len(seedProjects))
	return nil
}
