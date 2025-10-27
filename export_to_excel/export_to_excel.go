package main

import (
	"example-multi-module-workspace/db"
	"example-multi-module-workspace/projects"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2" // imported as 'excelize'
)

func main() {
	// Connect to MySQL
	database := db.Connect()
	defer database.Close()

	// Query your data
	data, err := projects.GetAll()
	if err != nil {
		log.Fatal("Fetch error:", err)
	}

	// Create Excel file
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	// Write header
	headers := []string{"ID", "Name", "Tech Stack"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Write rows
	rowIndex := 2
	for _, p := range data {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowIndex), p.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowIndex), p.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowIndex), p.TechStack)
		rowIndex++
	}

	// Save file
	if err := f.SaveAs("projects.xlsx"); err != nil {
		log.Fatal("Excel save error:", err)
	}

	fmt.Println("✅ Data exported successfully to projects.xlsx")
}
