package main

import (
	"fmt"

	"github.com/lk153/gsheet-go/lib"
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
	readRange := "'student'!A1:B4"
	values := srv.ReadSheet(spreadsheetID, readRange)
	for idx, row := range values {
		fmt.Println(idx, row)
	}
}
