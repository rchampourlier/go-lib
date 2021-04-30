package csv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ExtractCsvLineItems returns the values of the passed `line`.
//
// - `line` is a CSV string
func ExtractCsvLineItems(line string, sep string) []string {
	lineItems := strings.Split(line, sep)
	lineItems[0] = strings.TrimPrefix(lineItems[0], "\"")
	lineItems[len(lineItems)-1] = strings.TrimRight(lineItems[len(lineItems)-1], "\"\n\r")
	return lineItems
}

// ParseCsv parses a CSV file, returns a slice of rows with the values
// for each row and a map mapping the header string to the index of the
// value in the slice.
//
// Example:
//
// ```
// rows, headers := ParseCsvToRows("path/to/file.csv")
// result[3][headers["col1"]] // returns the value for "col1" column at row 3
// ```
//
func ParseCsvToRows(filepath string, sep string) (map[string]int, [][]string, error) {
	headers := make(map[string]int)
	rows := make([][]string, 0)

	file, err := os.Open(filepath)
	if err != nil {
		return headers, rows, fmt.Errorf("failed to open file: %v", err)
	}

	rowIdx := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rowItems := ExtractCsvLineItems(row, sep)
		if rowIdx == 0 {
			for i, rowItem := range rowItems {
				headers[rowItem] = i
			}
			rowIdx++
			continue
		}
		rows = append(rows, rowItems)
		rowIdx++
	}
	return headers, rows, nil
}

// ParseCsvToColumns parses a CSV file, returning a map of string
// slices representing each column.
//
// Example:
// result["col1"][3] returns the value for "col1" column at row 3.
func ParseCsvToColumns(filepath string, sep string) (map[string][]string, error) {
	columns := make(map[string][]string)
	var headers []string

	file, err := os.Open(filepath)
	if err != nil {
		return columns, fmt.Errorf("failed to open file: %v", err)
	}

	rowIdx := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rowItems := ExtractCsvLineItems(row, sep)
		if rowIdx == 0 {
			headers = make([]string, len(rowItems))
		}
		for colIdx, rowItem := range rowItems {
			if rowIdx == 0 {
				columns[rowItem] = make([]string, 0)
				headers[colIdx] = rowItem
			} else {
				colHeader := headers[colIdx]
				columns[colHeader] = append(columns[colHeader], rowItem)
			}
		}
		rowIdx++
	}

	return columns, nil
}
