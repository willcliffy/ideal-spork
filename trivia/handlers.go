package trivia

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/sheets/v4"
)

type TriviaHandler struct {
	NumPlayers  int
	Sheets      *sheets.SpreadsheetsService
	SheetID     string

	ScoreKeeper ScoreKeeper
}

func NewTriviaHandler(sheets *sheets.SpreadsheetsService, sheetId string, scoreKeeper ScoreKeeper) TriviaHandler {
	t := TriviaHandler{
		Sheets:  sheets,
		SheetID: sheetId,
		ScoreKeeper: scoreKeeper,
	}

	numPlayers, err := t.GetNumberOfPlayersFromStdIn()
	if err != nil {
		panic(err)
	}

	t.NumPlayers = numPlayers

	return t
}

func (t TriviaHandler) GetNumberOfPlayersFromStdIn() (int, error){
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
	r, err := strconv.ParseInt(text, 10, 64)
	return int(r), err
}

func (t TriviaHandler) WaitForAllSubmissions(round ROUND, numPlayers int) [][]interface{} {
	sheetRange := string(round) + "!A2:N15"
	getRequest := t.Sheets.Values.Get(t.SheetID, sheetRange)

	for {
		resp, err := getRequest.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}

		if len(resp.Values) != numPlayers {
			time.Sleep(3 * time.Second)
		} else {
			return resp.Values
		}
	}
}
