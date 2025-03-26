package main

import (
	"fmt"

	"github.com/lk153/gsheet-go/v2/lib"
)

func main() {
	runTest()
}

func runTest() {
	srv, err := lib.NewGsheetServiceV2()
	if err != nil {
		fmt.Println("Cannot connect Gsheet!")
		return
	}

	spreadsheetID := "1qqyC8O2ZpDg0BLb2APSc1EWkHo8heiooAaB7uPCzHuc"
	readRange := "'student'!A:D"
	values := srv.ReadSheet(spreadsheetID, readRange)
	for idx, row := range values {
		fmt.Println(idx, row)
	}

	data := [][]any{
		{5, "apple", 3.14, true},
		{6, "banana", 6.28, false},
		{7, "cherry", 9.42, true},
	}

	_, err = srv.Append(spreadsheetID, readRange, data)
	if err != nil {
		fmt.Println("Gsheet Append has error", err)
	}
}
