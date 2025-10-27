# Complete Migration Guide

## 📋 Overview

A complete database migration system has been created for this project. It includes:

- **SQL Migration Files** - Version-controlled database schema changes
- **Go Migration Functions** - Programmatic migration execution
- **CLI Tool** - Command-line interface for running migrations
- **Shell Script** - Easy-to-use bash wrapper
- **Seed Data** - 10 sample projects for testing

## 📁 Files Created

```
db/
├── migrations/
│   └── 001_create_projects_table.sql    # SQL migration with seed data
├── migrate.go                            # Go migration functions
├── cmd/migrate/main.go                   # CLI tool
├── run-migrations.sh                     # Shell script wrapper
└── MIGRATIONS.md                         # Detailed documentation

Root:
├── MIGRATION_SETUP.md                    # Quick start guide
└── MIGRATION_GUIDE.md                    # This file
```

## 🗄️ Database Schema

### Projects Table

Analyzed from the existing code (`projects/model.go` and `projects/repository.go`):

```sql
CREATE TABLE projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tech_stack VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

**Fields:**
- `id` - Auto-incrementing primary key
- `name` - Project name (required)
- `tech_stack` - Technology stack (required)
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

## 🌱 Seed Data

10 sample projects are automatically inserted:

| ID | Name | Tech Stack |
|----|------|-----------|
| 1 | E-Commerce Platform | Go, PostgreSQL, React |
| 2 | Real-time Chat Application | Go, WebSocket, Vue.js |
| 3 | Data Analytics Dashboard | Go, MySQL, Angular |
| 4 | Mobile API Gateway | Go, gRPC, Kubernetes |
| 5 | Content Management System | Go, MongoDB, Next.js |
| 6 | IoT Device Manager | Go, MQTT, Flutter |
| 7 | Microservices Framework | Go, Docker, Kubernetes |
| 8 | Machine Learning Pipeline | Go, Python, TensorFlow |
| 9 | Payment Processing System | Go, MySQL, React |
| 10 | Social Media Backend | Go, Redis, PostgreSQL |

## 🚀 Quick Start

### Method 1: Using Shell Script (Easiest)

```bash
cd db
chmod +x run-migrations.sh
./run-migrations.sh all
```

### Method 2: Using CLI Tool

```bash
cd db/cmd/migrate
go build -o migrate
./migrate migrate-and-seed
```

### Method 3: Using Go Code

```go
package main

import (
    "log"
    "example-multi-module-workspace/db"
)

func main() {
    database := db.Connect()
    defer database.Close()

    if err := db.RunMigrations(); err != nil {
        log.Fatal(err)
    }

    if err := db.SeedDatabase(); err != nil {
        log.Fatal(err)
    }
}
```

### Method 4: Direct SQL

```bash
mysql -u root -p testdb < db/migrations/001_create_projects_table.sql
```

## 📝 Available Commands

### Shell Script Commands

```bash
./run-migrations.sh migrate        # Run migrations only
./run-migrations.sh seed           # Seed database only
./run-migrations.sh all            # Run migrations and seed
./run-migrations.sh help           # Show help
```

### CLI Tool Commands

```bash
./migrate migrate                  # Run migrations
./migrate seed                     # Seed database
./migrate migrate-and-seed         # Run both
./migrate help                     # Show help
```

## 🔧 Configuration

Database connection is configured in `db/connect.go`:

```go
dsn := "root:root@tcp(localhost:3307)/testdb"
```

**To change connection details:**
1. Edit `db/connect.go`
2. Update the `dsn` variable
3. Rebuild the migration tool

## ✅ Verification

After running migrations, verify the setup:

```bash
# Connect to MySQL
mysql -u root -p testdb

# Check table exists
SHOW TABLES;

# Check data
SELECT * FROM projects;

# Count records
SELECT COUNT(*) FROM projects;
```

Expected output: 10 rows

## 🔄 How It Works

### Migration Process

1. **RunMigrations()** function:
   - Reads all `.sql` files from `db/migrations/`
   - Executes them in alphabetical order
   - Splits statements by semicolon
   - Logs each migration

2. **SeedDatabase()** function:
   - Checks if projects table has data
   - Skips if data exists (prevents duplicates)
   - Inserts 10 sample projects
   - Logs completion

### Safety Features

- ✅ Checks database connection before running
- ✅ Prevents duplicate seeding
- ✅ Comprehensive error handling
- ✅ Detailed logging
- ✅ Graceful failure messages

## 📚 Integration with Existing Code

The migration system integrates seamlessly:

```go
// Existing code in projects/repository.go
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
    
    return projects, nil
}

// Now works with seeded data!
```

## 🆕 Adding New Migrations

To add a new migration:

1. Create a new file: `db/migrations/002_your_migration.sql`
2. Write SQL statements (separate with semicolons)
3. Run migrations: `./run-migrations.sh migrate`

Example:

```sql
-- db/migrations/002_add_users_table.sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, email) VALUES
    ('admin', 'admin@example.com'),
    ('user1', 'user1@example.com');
```

## 🐛 Troubleshooting

### "database not connected"
- Ensure MySQL is running
- Check DSN in `db/connect.go`
- Verify database exists

### "permission denied" on shell script
```bash
chmod +x db/run-migrations.sh
```

### "table already exists"
- This is expected - migrations use `CREATE TABLE IF NOT EXISTS`
- Safe to run multiple times

### Seed data not inserted
- Check if table already has data
- Clear table: `TRUNCATE TABLE projects;`
- Re-run seed: `./run-migrations.sh seed`

## 📖 Documentation Files

- **MIGRATION_GUIDE.md** - This file (complete overview)
- **MIGRATION_SETUP.md** - Quick start guide
- **db/MIGRATIONS.md** - Detailed technical documentation

## 🎯 Next Steps

1. ✅ Ensure MySQL is running
2. ✅ Run migrations: `./run-migrations.sh all`
3. ✅ Verify data: `SELECT * FROM projects;`
4. ✅ Use in your application: `projects.GetAll()`
5. ✅ Add more migrations as needed

## 📞 Support

For issues or questions:
1. Check the troubleshooting section above
2. Review `db/MIGRATIONS.md` for detailed docs
3. Check database connection in `db/connect.go`
4. Verify MySQL is running and accessible

