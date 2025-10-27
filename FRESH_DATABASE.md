# 🔄 Fresh Database Command

## Overview

The `fresh` command drops the entire database and recreates it from scratch with all migrations and seed data. This is useful for:

- **Development:** Reset to a clean state
- **Testing:** Start with known data
- **Debugging:** Eliminate data corruption issues
- **Cleanup:** Remove old or test data

## ⚠️ Warning

**This command will DELETE all data in the database!**

Use with caution in production environments.

## Usage

### Basic Command

```bash
cd /home/manish/code/project_go/example-multi-module-workspace
chmod +x db/run-migrations.sh
./db/run-migrations.sh fresh
```

### What It Does

1. **Prompts for confirmation** - Asks you to type "yes" to confirm
2. **Drops the database** - Removes the entire `testdb` database
3. **Creates fresh database** - Creates a new empty `testdb`
4. **Runs migrations** - Creates all tables from migration files
5. **Seeds data** - Inserts 10 sample projects

### Step-by-Step Example

```bash
$ ./db/run-migrations.sh fresh

⚠️  This will DROP the entire database and recreate it fresh!
⚠️  All data will be lost!

Are you sure? Type 'yes' to confirm: yes

ℹ️  Dropping database 'testdb'...
✅ Database dropped

ℹ️  Creating fresh database 'testdb'...
✅ Fresh database created

ℹ️  Running migrations and seeding database...
ℹ️  Building migration tool...
✅ Migration tool built successfully

ℹ️  Executing migrations and seed...
✅ Migrations and seeding completed successfully!
```

## Comparison with Other Commands

| Command | Action | Data Loss |
|---------|--------|-----------|
| `migrate` | Run migrations only | No |
| `seed` | Insert seed data | No |
| `all` | Run migrations + seed | No (if table empty) |
| `fresh` | Drop DB + recreate + migrate + seed | **YES** |

## When to Use Each Command

### Use `migrate`
- When you've added new migration files
- When you want to update the schema
- When you don't want to lose existing data

### Use `seed`
- When you want to add sample data
- When the table is empty
- When you want to reset to sample data

### Use `all`
- First time setup
- When starting fresh development
- When you want migrations + seed in one command

### Use `fresh`
- When you want a completely clean database
- When you want to reset to initial state
- When you're debugging data issues
- When you want to start over

## Safety Features

### 1. Confirmation Prompt
```bash
Are you sure? Type 'yes' to confirm:
```
You must type exactly "yes" to proceed. Any other input cancels the operation.

### 2. Error Handling
If MySQL connection fails:
```
❌ Failed to drop database. Check MySQL connection.
```

### 3. Verification
After completion, verify the fresh database:
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"
# Expected output: 10
```

## Prerequisites

Before using `fresh`, ensure:

1. **MySQL is running:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"
   ```

2. **Correct credentials in `db/connect.go`:**
   ```go
   dsn := "root:root@tcp(localhost:3306)/testdb"
   ```

3. **Migrations directory exists:**
   ```bash
   ls -la db/migrations/
   ```

## Troubleshooting

### Issue: "Failed to drop database"

**Cause:** MySQL connection failed

**Solution:**
```bash
# Check MySQL is running
mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"

# Check credentials in db/connect.go
# Update if needed and rebuild migration tool
```

### Issue: "Failed to create database"

**Cause:** MySQL connection failed or permission issue

**Solution:**
```bash
# Verify MySQL user has CREATE DATABASE permission
mysql -u root -p"root" -h localhost --port 3306 -e "SHOW GRANTS FOR 'root'@'localhost';"

# Should include: GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost'
```

### Issue: Operation cancelled

**Cause:** You didn't type "yes" exactly

**Solution:** Run the command again and type "yes" when prompted

## Examples

### Example 1: Fresh Start for Development

```bash
# Start fresh development session
./db/run-migrations.sh fresh

# Verify
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT * FROM projects LIMIT 3;"
```

### Example 2: Reset After Testing

```bash
# After running tests, reset to clean state
./db/run-migrations.sh fresh

# Continue development
```

### Example 3: Debug Data Issues

```bash
# If data seems corrupted, start fresh
./db/run-migrations.sh fresh

# Verify everything works
./test-migration.sh
```

## Related Commands

### Reset Without Dropping Database

If you just want to clear data without dropping the database:

```bash
# Clear all data
mysql -u root -p"root" -h localhost --port 3306 testdb -e "TRUNCATE TABLE projects;"

# Re-seed
./db/run-migrations.sh seed
```

### Backup Before Fresh

If you want to backup data before running fresh:

```bash
# Backup current database
mysqldump -u root -p"root" -h localhost --port 3306 testdb > backup_$(date +%Y%m%d_%H%M%S).sql

# Now safe to run fresh
./db/run-migrations.sh fresh
```

### Restore from Backup

```bash
# Restore from backup
mysql -u root -p"root" -h localhost --port 3306 testdb < backup_20250101_120000.sql
```

## Best Practices

1. **Always confirm before running** - Read the warning carefully
2. **Backup important data** - Use mysqldump before fresh
3. **Use in development only** - Avoid in production
4. **Test after fresh** - Run test script to verify
5. **Document your data** - Keep notes of important data

## Quick Reference

```bash
# Run fresh database
./db/run-migrations.sh fresh

# Verify result
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"

# Expected: 10 projects
```

## Summary

The `fresh` command is a powerful tool for:
- ✅ Complete database reset
- ✅ Clean development environment
- ✅ Testing from known state
- ✅ Debugging data issues

**Use with caution!** Always confirm before proceeding.

---

**Status:** ✅ Ready to use
**Last Updated:** 2025-10-27

