package trivia

import (
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

type TriviaHandler struct {
	Sheets  *sheets.SpreadsheetsService
	SheetID string
}

func Start(spreadsheets *sheets.SpreadsheetsService, spreadsheetId string) {
	readRange := "Class Data!A2:E"

	resp, err := spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}
}

func (t TriviaHandler) roundScored() {
	// t.Sheets.Values.Get()
}
