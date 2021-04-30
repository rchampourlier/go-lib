package csv_test

import (
	"fmt"
	"testing"

	"golib/csv"
)

const CSV_TEST_FILE = "./csv_test_file.csv"
const CSV_SEP = ","

func TestParseCsvToRows(t *testing.T) {
	headers, rows, err := csv.ParseCsvToRows(CSV_TEST_FILE, CSV_SEP)
	expectedHeaders := []string{"col0", "col1", "col2"}
	expectedRows := [][]string{
		[]string{"r0c0", "r0c1", "r0c2"},
		[]string{"r1c0", "r1c1", "r1c2"},
	}

	if err != nil {
		t.Errorf("Failed to parse CSV file: %v", err)
	}

	if len(headers) != len(expectedHeaders) {
		t.Errorf("Expected %d headers, got %d", len(expectedHeaders), len(headers))
	}

	if len(rows) != len(expectedRows) {
		t.Errorf("Expected %d rows, got %d", len(expectedRows), len(rows))
	}

	for i, row := range rows {
		for j, item := range row {
			if item != expectedRows[i][j] {
				t.Errorf("Expected item at row %d and col %d to be `%s`, got `%s`", i, j, expectedRows[i][j], item)
			}
		}
	}
}

func TestParseCsvToColumns(t *testing.T) {
	columns, err := csv.ParseCsvToColumns(CSV_TEST_FILE, CSV_SEP)
	fmt.Println(columns)
	expectedHeaders := []string{"col0", "col1", "col2"}

	if err != nil {
		t.Errorf("Failed to parse CSV file: %v", err)
	}

	if len(columns) != 3 {
		t.Errorf("Expected 3 columns, got %d", len(columns))
	}

	for colHeader, col := range columns {
		if len(col) != 2 {
			t.Errorf("Expected column `%s` to have 2 items, got %d (%v)", colHeader, len(col), col)
		}
	}

	colIdx := 0
	for colHeader := range columns {
		expectedColumn := expectedHeaders[colIdx]
		if colHeader != expectedColumn {
			t.Errorf("Expected column `%s`, got `%s`", expectedColumn, colHeader)
		}
		colIdx++
	}

	if columns["col2"][1] != "r1c2" {
		t.Errorf("Expected item at `col2` and row 1 to be `%s`, got `%s`", "r1c2", columns["col2"][1])
	}
}
