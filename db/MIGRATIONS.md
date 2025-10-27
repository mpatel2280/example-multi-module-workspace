# Database Migrations

This document describes the database migration system for the project.

## Overview

The migration system provides a way to:
- Create and manage database schema changes
- Seed the database with initial data
- Maintain version control of database structure

## Migration Files

Migration files are located in `db/migrations/` and follow the naming convention:
```
NNN_description.sql
```

Where `NNN` is a zero-padded sequence number (e.g., `001`, `002`, etc.).

### Current Migrations

#### 001_create_projects_table.sql
Creates the `projects` table with the following schema:

```sql
CREATE TABLE projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tech_stack VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

**Seed Data:**
- 10 sample projects with various tech stacks
- Includes projects like E-Commerce Platform, Chat Application, Analytics Dashboard, etc.

## Usage

### Using the CLI Tool

The migration CLI tool is located at `db/cmd/migrate/main.go`.

#### Build the tool:
```bash
cd db/cmd/migrate
go build -o migrate
```

#### Run migrations:
```bash
./migrate migrate
```

#### Seed the database:
```bash
./migrate seed
```

#### Run migrations and seed together:
```bash
./migrate migrate-and-seed
```

### Using in Go Code

You can also call the migration functions directly from your Go code:

```go
package main

import (
    "log"
    "example-multi-module-workspace/db"
)

func main() {
    // Connect to database
    database := db.Connect()
    defer database.Close()

    // Run migrations
    if err := db.RunMigrations(); err != nil {
        log.Fatal(err)
    }

    // Seed database
    if err := db.SeedDatabase(); err != nil {
        log.Fatal(err)
    }
}
```

## Database Connection

The database connection is configured in `db/connect.go`:

```
DSN: root:root@tcp(localhost:3306)/testdb
```

To change the connection details, edit the `dsn` variable in `db/connect.go`.

## Seed Data

The seed data includes 10 sample projects:

1. E-Commerce Platform - Go, PostgreSQL, React
2. Real-time Chat Application - Go, WebSocket, Vue.js
3. Data Analytics Dashboard - Go, MySQL, Angular
4. Mobile API Gateway - Go, gRPC, Kubernetes
5. Content Management System - Go, MongoDB, Next.js
6. IoT Device Manager - Go, MQTT, Flutter
7. Microservices Framework - Go, Docker, Kubernetes
8. Machine Learning Pipeline - Go, Python, TensorFlow
9. Payment Processing System - Go, MySQL, React
10. Social Media Backend - Go, Redis, PostgreSQL

**Note:** The seed function checks if data already exists and skips seeding if the projects table is not empty.

## Adding New Migrations

To add a new migration:

1. Create a new SQL file in `db/migrations/` with the next sequence number
2. Write your SQL statements (separate multiple statements with semicolons)
3. Run the migration tool to apply the changes

Example:
```bash
# Create file: db/migrations/002_add_users_table.sql
cat > db/migrations/002_add_users_table.sql << 'EOF'
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
EOF

# Run the migration
./migrate migrate
```

## Troubleshooting

### Migration fails with "database not connected"
- Ensure MySQL is running
- Check the DSN in `db/connect.go`
- Verify the database exists

### Seed data not inserted
- Check if the projects table already has data
- The seed function skips if data exists to prevent duplicates
- Clear the table manually if you want to re-seed

### Permission denied when running migrate tool
```bash
chmod +x db/cmd/migrate/migrate
```

## Best Practices

1. **Always test migrations locally first** before running in production
2. **Keep migrations small and focused** on a single change
3. **Use descriptive names** for migration files
4. **Never modify existing migration files** - create new ones instead
5. **Document schema changes** in the migration file comments
6. **Backup your database** before running migrations in production

