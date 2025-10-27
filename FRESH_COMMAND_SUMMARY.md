# 🔄 Fresh Command - Feature Summary

## What's New

A new `fresh` command has been added to the migration system that allows you to drop and recreate the database from scratch.

## Quick Start

```bash
./db/run-migrations.sh fresh
```

## What It Does

The `fresh` command performs these steps:

1. **Prompts for confirmation** - Asks you to type "yes" to confirm
2. **Drops the database** - Removes the entire `testdb` database
3. **Creates fresh database** - Creates a new empty `testdb`
4. **Runs migrations** - Creates all tables from migration files
5. **Seeds data** - Inserts 10 sample projects

## Use Cases

### Development
```bash
# Start a fresh development session
./db/run-migrations.sh fresh
```

### Testing
```bash
# Reset to known state before running tests
./db/run-migrations.sh fresh
./run-tests.sh
```

### Debugging
```bash
# If data seems corrupted, start fresh
./db/run-migrations.sh fresh
```

### Cleanup
```bash
# Remove old test data
./db/run-migrations.sh fresh
```

## Safety Features

### 1. Confirmation Required
```bash
⚠️  This will DROP the entire database and recreate it fresh!
⚠️  All data will be lost!

Are you sure? Type 'yes' to confirm:
```

You must type exactly "yes" to proceed.

### 2. Error Handling
If MySQL connection fails, the command stops with an error message.

### 3. Verification
After completion, verify the fresh database:
```bash
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"
# Expected: 10
```

## Command Comparison

| Command | Purpose | Data Loss |
|---------|---------|-----------|
| `migrate` | Run migrations | No |
| `seed` | Insert seed data | No |
| `all` | Migrate + seed | No |
| `fresh` | Drop + recreate + migrate + seed | **YES** |

## Examples

### Example 1: Fresh Start
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
✅ Migrations and seeding completed successfully!
```

### Example 2: Cancel Operation
```bash
$ ./db/run-migrations.sh fresh

⚠️  This will DROP the entire database and recreate it fresh!
⚠️  All data will be lost!

Are you sure? Type 'yes' to confirm: no

ℹ️  Operation cancelled
```

### Example 3: Verify Result
```bash
$ ./db/run-migrations.sh fresh
# ... (runs fresh command)

$ mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"
COUNT(*)
10
```

## Prerequisites

Before using `fresh`:

1. **MySQL is running:**
   ```bash
   mysql -u root -p"root" -h localhost --port 3306 -e "SELECT 1;"
   ```

2. **Correct credentials in `db/connect.go`**

3. **Migrations directory exists:**
   ```bash
   ls -la db/migrations/
   ```

## Troubleshooting

### "Failed to drop database"
- Check MySQL is running
- Verify credentials in `db/connect.go`
- Check MySQL user has DROP DATABASE permission

### "Failed to create database"
- Check MySQL is running
- Verify MySQL user has CREATE DATABASE permission
- Check disk space

### Operation cancelled
- You didn't type "yes" exactly
- Run the command again

## Related Commands

### Reset Without Dropping
```bash
# Clear data but keep database
mysql -u root -p"root" -h localhost --port 3306 testdb -e "TRUNCATE TABLE projects;"

# Re-seed
./db/run-migrations.sh seed
```

### Backup Before Fresh
```bash
# Backup current database
mysqldump -u root -p"root" -h localhost --port 3306 testdb > backup.sql

# Now safe to run fresh
./db/run-migrations.sh fresh
```

### Restore from Backup
```bash
# Restore from backup
mysql -u root -p"root" -h localhost --port 3306 testdb < backup.sql
```

## Documentation

For detailed information, see:
- `FRESH_DATABASE.md` - Complete fresh command documentation
- `QUICK_REFERENCE.md` - Quick command reference
- `MIGRATION_GUIDE.md` - Complete migration guide

## Implementation Details

### File Modified
- `db/run-migrations.sh` - Added `run_fresh()` function and `fresh` command

### New Function
```bash
run_fresh() {
    # 1. Prompt for confirmation
    # 2. Drop database
    # 3. Create fresh database
    # 4. Run migrations and seed
}
```

### Safety Checks
- Confirmation prompt (must type "yes")
- Error handling for MySQL operations
- Graceful failure messages

## Best Practices

1. **Always confirm** - Read the warning carefully
2. **Backup first** - Use mysqldump before fresh
3. **Use in development** - Avoid in production
4. **Test after** - Run test script to verify
5. **Document data** - Keep notes of important data

## Quick Reference

```bash
# Run fresh database
./db/run-migrations.sh fresh

# Verify
mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;"

# Expected: 10
```

## Summary

✅ **New Feature:** `fresh` command
✅ **Purpose:** Drop and recreate database
✅ **Safety:** Confirmation required
✅ **Use Cases:** Development, testing, debugging
✅ **Status:** Ready to use

---

**Added:** 2025-10-27
**Status:** ✅ Complete

