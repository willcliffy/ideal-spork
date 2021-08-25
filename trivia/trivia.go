package trivia

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/sheets/v4"
)

type TriviaHandler struct {
	Sheets  *sheets.SpreadsheetsService
	SheetID string
}

func Start(spreadsheets *sheets.SpreadsheetsService, spreadsheetId string) {
	t := TriviaHandler{
		Sheets: spreadsheets,
		SheetID: spreadsheetId,
	}

	t.RoundOne()
	t.RoundTwo()
	t.RoundThree()
	t.WaitForWagers()
	t.RoundFinal()
}

func (t TriviaHandler) WaitForAllSubmissions(round ROUND) [][]interface{} {
	sheetRange := string(round) + "!A2:N15"
	getRequest := t.Sheets.Values.Get(t.SheetID, sheetRange)

	for {
		resp, err := getRequest.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
	
		if len(resp.Values) == 0 {
			fmt.Println("No data found.")
			time.Sleep(3 * time.Second)
		} else {
			return resp.Values
		}
	}
}

func (t TriviaHandler) RoundOne() {
	res := t.WaitForAllSubmissions(RoundOne)

	for i, v := range res {
		log.Printf("%v: %v", i, v)
	}
}

func (t TriviaHandler) RoundTwo() {
	res := t.WaitForAllSubmissions(RoundTwo)
	for i, v := range res {
		log.Printf("%v: %v", i, v)
	}
}

func (t TriviaHandler) RoundThree() {
	res := t.WaitForAllSubmissions(RoundThree)
	for i, v := range res {
		log.Printf("%v: %v", i, v)
	}
}

func (t TriviaHandler) WaitForWagers() {

}

func (t TriviaHandler) RoundFinal() {
	res := t.WaitForAllSubmissions(RoundFinal)
	for i, v := range res {
		log.Printf("%v: %v", i, v)
	}
}
