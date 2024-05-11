package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/tealeg/xlsx"
)

// Data represents the structure of data you expect from the Excel sheet
type Data struct {
	Item_name string
	Items     string
}

func main() {
	// Load Excel file
	xlFile, err := xlsx.OpenFile("cmd/app/items_list.xlsx")
	if err != nil {
		log.Fatalf("Error opening Excel file: %v", err)
	}

	// Assuming the first sheet contains the data
	sheet := xlFile.Sheets[0]

	// Parse Excel data into a slice of Data
	var data []Data
	for _, row := range sheet.Rows[1:] { // Assuming the first row is the header
		d := Data{
			Item_name: row.Cells[0].String(),
			Items:     row.Cells[1].String(),
		}
		data = append(data, d)
	}

	// Define a handler to serve HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("cmd/app/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
