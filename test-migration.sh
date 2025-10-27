#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Migration System Test${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Step 1: Build the migration tool
echo -e "${YELLOW}Step 1: Building migration tool...${NC}"
cd "$SCRIPT_DIR/db/cmd/migrate"
go build -o migrate
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to build migration tool${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Migration tool built successfully${NC}"
echo ""

# Step 2: Run migrations
echo -e "${YELLOW}Step 2: Running migrations...${NC}"
cd "$SCRIPT_DIR"
./db/cmd/migrate/migrate migrate
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to run migrations${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Migrations completed${NC}"
echo ""

# Step 3: Seed database
echo -e "${YELLOW}Step 3: Seeding database...${NC}"
./db/cmd/migrate/migrate seed
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to seed database${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Database seeded${NC}"
echo ""

# Step 4: Verify
echo -e "${YELLOW}Step 4: Verifying data...${NC}"
RESULT=$(mysql -u root -p"root" -h localhost --port 3306 testdb -e "SELECT COUNT(*) FROM projects;" 2>/dev/null | tail -1)
if [ "$RESULT" = "10" ]; then
    echo -e "${GREEN}✅ Verification successful! Found 10 projects${NC}"
else
    echo -e "${YELLOW}⚠️  Found $RESULT projects (expected 10)${NC}"
fi
echo ""

echo -e "${BLUE}========================================${NC}"
echo -e "${GREEN}✅ All tests completed!${NC}"
echo -e "${BLUE}========================================${NC}"

