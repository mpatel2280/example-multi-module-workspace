package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"example-multi-module-workspace/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Define command-line flags
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	seedCmd := flag.NewFlagSet("seed", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "migrate":
		migrateCmd.Parse(os.Args[2:])
		runMigrations()

	case "seed":
		seedCmd.Parse(os.Args[2:])
		runSeed()

	case "migrate-and-seed":
		migrateCmd.Parse(os.Args[2:])
		runMigrationsAndSeed()

	case "fresh":
		runFresh()

	case "refresh":
		runReFresh()

	case "help":
		printUsage()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func runMigrations() {
	fmt.Println("🔄 Running database migrations...")

	// Connect to database
	database := db.Connect()
	defer database.Close()

	// Run migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	fmt.Println("✅ All migrations completed successfully!")
}

func runSeed() {
	fmt.Println("🌱 Seeding database...")

	// Connect to database
	database := db.Connect()
	defer database.Close()

	// Seed database
	if err := db.SeedDatabase(); err != nil {
		log.Fatalf("❌ Seeding failed: %v", err)
	}

	fmt.Println("✅ Database seeded successfully!")
}

func runMigrationsAndSeed() {
	fmt.Println("🔄 Running migrations and seeding database...")

	// Connect to database
	database := db.Connect()
	defer database.Close()

	// Run migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	// Seed database
	if err := db.SeedDatabase(); err != nil {
		log.Fatalf("❌ Seeding failed: %v", err)
	}

	fmt.Println("✅ Migrations and seeding completed successfully!")
}

func runFresh() {
	fmt.Println("🔄 Fresh migration: Drop database and recreate with seed data...")
	fmt.Println("")
	fmt.Println("⚠️  WARNING: This will DROP the entire database and recreate it fresh!")
	fmt.Println("⚠️  All data will be lost!")
	fmt.Println("")

	// Prompt for confirmation
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure? Type 'yes' to confirm: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input != "yes" {
		fmt.Println("ℹ️  Operation cancelled")
		return
	}

	// Connect to MySQL without specifying database
	fmt.Println("ℹ️  Connecting to MySQL...")
	dsn := "root:root@tcp(localhost:3306)/"
	mysqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ MySQL connection failed: %v", err)
	}
	defer mysqlDB.Close()

	if err := mysqlDB.Ping(); err != nil {
		log.Fatalf("❌ MySQL ping failed: %v", err)
	}

	// Drop database
	fmt.Println("ℹ️  Dropping database 'testdb'...")
	_, err = mysqlDB.Exec("DROP DATABASE IF EXISTS testdb")
	if err != nil {
		log.Fatalf("❌ Failed to drop database: %v", err)
	}
	fmt.Println("✅ Database dropped")

	// Create database
	fmt.Println("ℹ️  Creating fresh database 'testdb'...")
	_, err = mysqlDB.Exec("CREATE DATABASE IF NOT EXISTS testdb")
	if err != nil {
		log.Fatalf("❌ Failed to create database: %v", err)
	}
	fmt.Println("✅ Fresh database created")

	fmt.Println("")

	// Now run migrations and seed
	fmt.Println("ℹ️  Running migrations and seeding database...")
	database := db.Connect()
	defer database.Close()

	// Run migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	// Seed database
	if err := db.SeedDatabase(); err != nil {
		log.Fatalf("❌ Seeding failed: %v", err)
	}

	fmt.Println("✅ Fresh migration completed successfully!")
	fmt.Println("✅ Database dropped, recreated, and seeded with sample data!")
}

func runReFresh() {
	fmt.Println("🔄 Fresh migration: Drop database and recreate with seed data...")
	fmt.Println("")
	fmt.Println("⚠️  WARNING: This will DROP the entire database and recreate it fresh!")
	fmt.Println("⚠️  All data will be lost!")
	fmt.Println("")

	// Prompt for confirmation
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure? Type 'yes' to confirm: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input != "yes" {
		fmt.Println("ℹ️  Operation cancelled")
		return
	}

	// Connect to database
	database := db.Connect()
	defer database.Close()

	// Run migrations
	if err := db.RunRefreshMigrations(); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	fmt.Println("✅ Fresh migration completed successfully!")
	fmt.Println("✅ Database dropped, recreated, and seeded with sample data!")
}

func printUsage() {
	fmt.Println(`Database Migration Tool

Usage:
  migrate <command> [options]

Commands:
  migrate              Run all pending migrations
  seed                 Seed the database with sample data
  migrate-and-seed     Run migrations and seed in one command
  fresh                Drop database and recreate it fresh with seed data
  refresh              Drop database and recreate it fresh with seed data (Using Single SQL Script file)
  help                 Show this help message

Examples:
  migrate migrate
  migrate seed
  migrate migrate-and-seed
  migrate fresh
  migrate refresh
  migrate help

Environment Variables:
  Database connection is configured in db/connect.go
  Default: root:root@tcp(localhost:3306)/testdb

Note:
  The 'fresh' command will DROP the entire database and recreate it.
  Use with caution! You will be prompted for confirmation.`)
}
