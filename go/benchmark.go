package main

import (
	"strconv"
	"fmt"
	"log"
	"time"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	rows := 100000
	cols := 50

	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
		for j := range data[i] {
			data[i][j] = i + j
		}
	}


	beginTimestamp := time.Now()
	ss := spreadsheet.New()
	// add a single sheet
	sheet := ss.AddSheet()

	// rows
	for rowIndex := 0; rowIndex < rows; rowIndex++ {
		row := sheet.AddRow()

		for col := 0; col < cols; col++ {
			cell := row.AddCell()
			cell.SetString(strconv.Itoa(data[rowIndex][col]))
		}
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("output/go.xlsx")
	endTimestamp := time.Now()
	elapsed := endTimestamp.Sub(beginTimestamp)

	fmt.Printf("Go: Writing 10000x50 cells of data takes %v seconds", elapsed)
}
