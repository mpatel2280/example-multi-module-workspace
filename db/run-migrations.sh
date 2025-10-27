#!/bin/bash

# Database Migration Runner Script
# This script provides an easy way to run database migrations and seed data

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the directory where this script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# Function to show usage
show_usage() {
    cat << EOF
${BLUE}Database Migration Runner${NC}

Usage: ./run-migrations.sh [command]

Commands:
  migrate              Run all pending migrations
  seed                 Seed the database with sample data
  all                  Run migrations and seed (default)
  help                 Show this help message

Examples:
  ./run-migrations.sh migrate
  ./run-migrations.sh seed
  ./run-migrations.sh all
  ./run-migrations.sh help

Environment:
  Database connection configured in: db/connect.go
  Migrations directory: db/migrations/
  Seed data: 10 sample projects

EOF
}

# Function to run migrations
run_migrations() {
    print_info "Running database migrations..."

    cd "$SCRIPT_DIR/cmd/migrate"

    # Build the migration tool
    print_info "Building migration tool..."
    go build -o migrate

    if [ ! -f migrate ]; then
        print_error "Failed to build migration tool"
        exit 1
    fi

    print_success "Migration tool built successfully"

    # Run migrations from the project root directory
    print_info "Executing migrations..."
    cd "$PROJECT_ROOT"
    "$SCRIPT_DIR/cmd/migrate/migrate" migrate

    print_success "Migrations completed successfully!"
}

# Function to seed database
run_seed() {
    print_info "Seeding database..."

    cd "$SCRIPT_DIR/cmd/migrate"

    # Build the migration tool if not already built
    if [ ! -f migrate ]; then
        print_info "Building migration tool..."
        go build -o migrate
    fi

    # Run seed from the project root directory
    print_info "Inserting seed data..."
    cd "$PROJECT_ROOT"
    "$SCRIPT_DIR/cmd/migrate/migrate" seed

    print_success "Database seeded successfully!"
}

# Function to run both migrations and seed
run_all() {
    print_info "Running migrations and seeding database..."

    cd "$SCRIPT_DIR/cmd/migrate"

    # Build the migration tool
    print_info "Building migration tool..."
    go build -o migrate

    if [ ! -f migrate ]; then
        print_error "Failed to build migration tool"
        exit 1
    fi

    print_success "Migration tool built successfully"

    # Run migrations and seed from the project root directory
    print_info "Executing migrations and seed..."
    cd "$PROJECT_ROOT"
    "$SCRIPT_DIR/cmd/migrate/migrate" migrate-and-seed

    print_success "Migrations and seeding completed successfully!"
}

# Main script logic
main() {
    local command="${1:-all}"
    
    case "$command" in
        migrate)
            run_migrations
            ;;
        seed)
            run_seed
            ;;
        all)
            run_all
            ;;
        help)
            show_usage
            ;;
        *)
            print_error "Unknown command: $command"
            echo ""
            show_usage
            exit 1
            ;;
    esac
}

# Run main function
main "$@"

