# 🔧 Migration System - Troubleshooting Guide

## Common Issues and Solutions

### Issue 1: "failed to read migrations directory: no such file or directory"

**Cause:** The migration tool cannot find the `db/migrations` directory.

**Solutions:**

1. **Run from project root:**
   ```bash
   cd /home/manish/code/project_go/example-multi-module-workspace
   ./db/cmd/migrate/migrate migrate
   ```

2. **Use the shell script (recommended):**
   ```bash
   cd /home/manish/code/project_go/example-multi-module-workspace
   chmod +x db/run-migrations.sh
   ./db/run-migrations.sh all
   ```

3. **Check migrations directory exists:**
   ```bash
   ls -la db/migrations/
   # Should show: 001_create_projects_table.sql
   ```

### Issue 2: "database not connected"

**Cause:** The database connection failed.

**Solutions:**

1. **Verify MySQL is running:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"
   ```

2. **Check connection details in `db/connect.go`:**
   ```go
   dsn := "root:root@tcp(localhost:3306)/testdb"
   ```

3. **Verify database exists:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 -e "SHOW DATABASES;" | grep testdb
   ```

4. **Create database if missing:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 -e "CREATE DATABASE IF NOT EXISTS testdb;"
   ```

### Issue 3: "permission denied" on shell script

**Cause:** Shell script doesn't have execute permissions.

**Solution:**
```bash
chmod +x db/run-migrations.sh
./db/run-migrations.sh all
```

### Issue 4: "table already exists"

**Cause:** The projects table was already created.

**Solution:** This is normal! Migrations use `CREATE TABLE IF NOT EXISTS`, so it's safe to run multiple times.

If you want to reset:
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb -e "DROP TABLE IF EXISTS projects;"
./db/run-migrations.sh migrate
```

### Issue 5: Seed data not inserted

**Cause:** Table already has data, or seed failed.

**Solutions:**

1. **Check if data exists:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"
   ```

2. **Clear and reseed:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 testdb -e "TRUNCATE TABLE projects;"
   ./db/run-migrations.sh seed
   ```

3. **Check seed output:**
   ```bash
   ./db/cmd/migrate/migrate seed
   # Look for: "✅ Seeded 10 projects into the database"
   ```

### Issue 6: "go build" fails

**Cause:** Go dependencies not installed or version mismatch.

**Solutions:**

1. **Download dependencies:**
   ```bash
   cd db
   go mod download
   go mod tidy
   ```

2. **Check Go version:**
   ```bash
   go version
   # Should be Go 1.16 or later
   ```

3. **Rebuild:**
   ```bash
   cd db/cmd/migrate
   go build -o migrate
   ```

### Issue 7: MySQL connection refused

**Cause:** MySQL not running or wrong port.

**Solutions:**

1. **Check MySQL status:**
   ```bash
   # On Linux/Mac
   sudo systemctl status mysql
   
   # Or try to connect
   mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"
   ```

2. **Verify port (default is 3306):**
   ```bash
   # Check what port MySQL is listening on
   sudo netstat -tlnp | grep mysql
   ```

3. **Update connection if needed:**
   - Edit `db/connect.go`
   - Change the `dsn` variable
   - Rebuild the migration tool

### Issue 8: "Access denied for user 'root'@'localhost'"

**Cause:** Wrong password or user doesn't exist.

**Solutions:**

1. **Verify credentials:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"
   ```

2. **Update credentials in `db/connect.go`:**
   ```go
   dsn := "username:password@tcp(localhost:3306)/testdb"
   ```

3. **Reset MySQL root password (if needed):**
   ```bash
   # Consult MySQL documentation for your OS
   ```

### Issue 9: "no such file or directory" when running from different location

**Cause:** Working directory is not the project root.

**Solution:** Always run from project root:
```bash
cd /home/manish/code/project_go/example-multi-module-workspace
./db/run-migrations.sh all
```

Or use absolute paths:
```bash
/home/manish/code/project_go/example-multi-module-workspace/db/run-migrations.sh all
```

### Issue 10: Migrations run but no data appears

**Cause:** Seed function skipped because table already has data.

**Solution:**
```bash
# Check current data
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT * FROM projects;"

# If empty but count shows data, check for issues
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"

# Clear and reseed
mysql -u root -p"root" -h localhost --port 3306 testdb -e "TRUNCATE TABLE projects;"
./db/cmd/migrate/migrate seed
```

## Verification Steps

### Step 1: Check MySQL Connection
```bash
mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"
# Expected: 1
```

### Step 2: Check Database Exists
```bash
mysql -u root -p"root" -h localhost --port 3306 -e "SHOW DATABASES;" | grep testdb
# Expected: testdb
```

### Step 3: Check Table Exists
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SHOW TABLES;"
# Expected: projects
```

### Step 4: Check Data
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT * FROM projects;"
# Expected: 10 rows
```

### Step 5: Check Count
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"
# Expected: 10
```

## Quick Test Script

Run the included test script:
```bash
chmod +x test-migration.sh
./test-migration.sh
```

This will:
1. Build the migration tool
2. Run migrations
3. Seed the database
4. Verify the data

## Debug Mode

To see detailed output:

1. **Run migration with verbose output:**
   ```bash
   cd /home/manish/code/project_go/example-multi-module-workspace
   ./db/cmd/migrate/migrate migrate
   ```

2. **Check logs:**
   - Look for "✅" messages indicating success
   - Look for error messages indicating failures

3. **Manual SQL execution:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 testdb < db/migrations/001_create_projects_table.sql
   ```

## Getting Help

1. **Check documentation:**
   - `README_MIGRATIONS.md` - Quick reference
   - `MIGRATION_GUIDE.md` - Complete guide
   - `db/MIGRATIONS.md` - Technical details

2. **Verify setup:**
   - MySQL running and accessible
   - Database `testdb` exists
   - Correct credentials in `db/connect.go`

3. **Run test script:**
   ```bash
   ./test-migration.sh
   ```

## Common Configuration

**Default Configuration:**
- Host: localhost
- Port: 3306
- User: root
- Password: root
- Database: testdb

**To change:**
1. Edit `db/connect.go`
2. Update the `dsn` variable
3. Rebuild: `cd db/cmd/migrate && go build -o migrate`

## Still Having Issues?

1. Verify MySQL is running
2. Check database connection: `mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"`
3. Verify migrations directory exists: `ls -la db/migrations/`
4. Run from project root directory
5. Check file permissions: `chmod +x db/run-migrations.sh`
6. Review error messages carefully
7. Check `db/connect.go` for correct credentials

---

**Last Updated:** 2025-10-27
**Status:** Complete

