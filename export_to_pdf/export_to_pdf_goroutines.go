package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"example-multi-module-workspace/db"
	"example-multi-module-workspace/projects"

	"github.com/phpdave11/gofpdf"
)

func ExportToPDF(filename string, title string, data interface{}) error {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, title)
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)

	// Headers
	if v.Len() > 0 {
		elemType := v.Index(0).Type()
		for i := 0; i < elemType.NumField(); i++ {
			pdf.CellFormat(40, 10, elemType.Field(i).Name, "1", 0, "", false, 0, "")
		}
		pdf.Ln(10)
	}

	// Rows
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		for j := 0; j < elem.NumField(); j++ {
			val := fmt.Sprintf("%v", elem.Field(j).Interface())
			pdf.CellFormat(40, 10, val, "1", 0, "", false, 0, "")
		}
		pdf.Ln(10)
	}

	return pdf.OutputFileAndClose(filename)
}

// chunkSlice splits a slice into chunks of given size
func chunkSlice[T any](data []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}

func main() {
	dbConn := db.Connect()
	defer dbConn.Close()

	// 1️⃣ Fetch all data
	data, err := projects.GetAll()
	if err != nil {
		log.Fatal("Fetch error:", err)
	}

	// 2️⃣ Split into chunks (e.g. 100 per PDF)
	chunks := chunkSlice(data, 4)

	// 3️⃣ Use WaitGroup + channels to run concurrently
	var wg sync.WaitGroup
	results := make(chan string, len(chunks))
	errors := make(chan error, len(chunks))

	for i, chunk := range chunks {
		wg.Add(1)
		go func(i int, chunk []projects.Project) {
			defer wg.Done()
			filename := fmt.Sprintf("projects_report_part_%d.pdf", i+1)
			if err := ExportToPDF(filename, fmt.Sprintf("Projects Report Part %d", i+1), chunk); err != nil {
				errors <- fmt.Errorf("error in chunk %d: %v", i+1, err)
				return
			}
			results <- filename
		}(i, chunk)
	}

	// 4️⃣ Wait for completion
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	// 5️⃣ Collect results
	for {
		select {
		case file, ok := <-results:
			if !ok {
				results = nil
			} else {
				fmt.Println("✅ Generated:", file)
			}
		case err, ok := <-errors:
			if ok {
				fmt.Println("❌", err)
			} else {
				errors = nil
			}
		}
		if results == nil && errors == nil {
			break
		}
	}

	fmt.Println("🎉 All PDF chunks generated successfully!")
}
