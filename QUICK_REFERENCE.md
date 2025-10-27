# ⚡ Migration System - Quick Reference Card

## 🚀 One-Liner to Run Everything

```bash
cd /home/manish/code/project_go/example-multi-module-workspace && chmod +x db/run-migrations.sh && ./db/run-migrations.sh all
```

## 📋 All Commands

### Using Shell Script (Easiest)
```bash
cd db
chmod +x run-migrations.sh

./run-migrations.sh migrate        # Run migrations only
./run-migrations.sh seed           # Seed database only
./run-migrations.sh all            # Run both (default)
./run-migrations.sh help           # Show help
```

### Using CLI Tool
```bash
cd db/cmd/migrate
go build -o migrate

cd ../..  # Back to project root
./db/cmd/migrate/migrate migrate                  # Run migrations
./db/cmd/migrate/migrate seed                     # Seed database
./db/cmd/migrate/migrate migrate-and-seed         # Run both
./db/cmd/migrate/migrate help                     # Show help
```

### Using Go Code
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

### Using Direct SQL
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb < db/migrations/001_create_projects_table.sql
```

## ✅ Verification Commands

```bash
# Check table exists
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SHOW TABLES;"

# Check data
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT * FROM projects;"

# Count records (should be 10)
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"

# Check specific project
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT * FROM projects WHERE id = 1;"
```

## 🔧 Configuration

**File:** `db/connect.go`

**Current Settings:**
```
Host: localhost
Port: 3306
User: root
Password: root
Database: testdb
```

**To Change:**
1. Edit `db/connect.go`
2. Update the `dsn` variable
3. Rebuild: `cd db/cmd/migrate && go build -o migrate`

## 📁 File Locations

```
db/
├── migrations/
│   └── 001_create_projects_table.sql    ← SQL schema + seed data
├── migrate.go                            ← Migration functions
├── cmd/migrate/main.go                   ← CLI tool
├── run-migrations.sh                     ← Shell script wrapper
└── connect.go                            ← Database connection

Root:
├── README_MIGRATIONS.md                  ← Quick start
├── MIGRATION_GUIDE.md                    ← Complete guide
├── MIGRATION_SETUP.md                    ← Setup details
├── TROUBLESHOOTING.md                    ← Problem solving
├── QUICK_REFERENCE.md                    ← This file
└── test-migration.sh                     ← Test script
```

## 🌱 Seed Data (10 Projects)

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

## 🗄️ Database Schema

```sql
CREATE TABLE projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tech_stack VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## 🆕 Adding New Migrations

1. Create file: `db/migrations/002_your_migration.sql`
2. Write SQL (separate statements with semicolons)
3. Run: `./db/run-migrations.sh migrate`

Example:
```sql
-- db/migrations/002_add_users_table.sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE
);
```

## 🐛 Quick Troubleshooting

| Problem | Solution |
|---------|----------|
| "no such file or directory" | Run from project root: `cd /home/manish/code/project_go/example-multi-module-workspace` |
| "database not connected" | Check MySQL is running: `mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"` |
| "permission denied" | `chmod +x db/run-migrations.sh` |
| "table already exists" | Normal - safe to run multiple times |
| Seed data not inserted | Clear table: `mysql -u root -p"root" -h localhost --port 3306 testdb -e "TRUNCATE TABLE projects;"` |

## 📚 Documentation Map

- **2 min read:** `README_MIGRATIONS.md`
- **5 min read:** `MIGRATION_SETUP.md`
- **15 min read:** `MIGRATION_GUIDE.md`
- **10 min read:** `db/MIGRATIONS.md`
- **5 min read:** `TROUBLESHOOTING.md`
- **3 min read:** `QUICK_REFERENCE.md` (this file)

## 🎯 Typical Workflow

```bash
# 1. Navigate to project
cd /home/manish/code/project_go/example-multi-module-workspace

# 2. Make script executable
chmod +x db/run-migrations.sh

# 3. Run migrations and seed
./db/run-migrations.sh all

# 4. Verify
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"

# 5. Use in your code
# projects.GetAll() now returns 10 projects
```

## 🧪 Test Script

```bash
chmod +x test-migration.sh
./test-migration.sh
```

Runs:
1. Build migration tool
2. Run migrations
3. Seed database
4. Verify data

## 💡 Pro Tips

1. **Always run from project root** to avoid path issues
2. **Use shell script** for easiest experience
3. **Check MySQL is running** before running migrations
4. **Migrations are idempotent** - safe to run multiple times
5. **Seed only runs if table is empty** - prevents duplicates

## 🔗 Integration with Existing Code

```go
// In projects/repository.go
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
projects, _ := GetAll()
// Returns 10 projects
```

## 📞 Need Help?

1. Check `TROUBLESHOOTING.md` for common issues
2. Read `MIGRATION_GUIDE.md` for complete details
3. Run `test-migration.sh` to verify setup
4. Check MySQL connection: `mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"`

---

**Status:** ✅ Ready to use
**Last Updated:** 2025-10-27

