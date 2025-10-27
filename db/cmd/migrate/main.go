package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"example-multi-module-workspace/db"
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

func printUsage() {
	fmt.Println(`
Database Migration Tool

Usage:
  migrate <command> [options]

Commands:
  migrate              Run all pending migrations
  seed                 Seed the database with sample data
  migrate-and-seed     Run migrations and seed in one command
  help                 Show this help message

Examples:
  migrate migrate
  migrate seed
  migrate migrate-and-seed

Environment Variables:
  Database connection is configured in db/connect.go
  Default: root:root@tcp(localhost:3306)/testdb
`)
}
