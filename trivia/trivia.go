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

func Start(spreadsheets *sheets.SpreadsheetsService, spreadsheetId string) {
	t := TriviaHandler{
		Sheets: spreadsheets,
		SheetID: spreadsheetId,
	}

	t.Scores      = make(map[ROUND]map[string]int)
	t.Cumulative  = make(map[string]int)
	t.Wagers      = make(map[string]int)
	t.FinalScores = make(map[string]int)

	numPlayers, err := t.GetNumberOfPlayersFromStdIn()
	if err != nil {
		panic(err)
	}

	t.NumPlayers = numPlayers

	t.RoundOne()
	t.RoundTwo()
	t.RoundThree()
	t.WaitForWagers()
	t.RoundFinal()
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

func (t TriviaHandler) RoundOne() {
	res := t.WaitForAllSubmissions(RoundOne, t.NumPlayers)

	scores := make(map[string]int)
	for _, v := range res {
		// todo - error handling and fuzzy matching on emails and team names.
		//name  := v[3].(string)
		email := v[1].(string)

		scoreStr := strings.Split(v[2].(string), " / ")[0]
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			panic(err)
		}

		scores[email] = score
	}

	t.AddScoresToScoreboard(RoundOne, scores)
}

func (t TriviaHandler) RoundTwo() {
	res := t.WaitForAllSubmissions(RoundTwo, t.NumPlayers)

	for _, v := range res {
		log.Printf("name: %v, email: %v, score: %v", v[3], v[1], v[2])
	}
}

func (t TriviaHandler) RoundThree() {
	res := t.WaitForAllSubmissions(RoundThree, t.NumPlayers)

	for _, v := range res {
		log.Printf("name: %v, email: %v, \n\t21: %v, \n\t22: %v, \n\t23: %v, \n\t24: %v, \n\t25: %v", v[3], v[1], v[4], v[5], v[6], v[7], v[8])
	}
}

func (t TriviaHandler) WaitForWagers() {

}

func (t TriviaHandler) RoundFinal() {
	res := t.WaitForAllSubmissions(RoundFinal, t.NumPlayers)

	for _, v := range res {
		log.Printf("name: %v, email: %v, final answer: %v", v[2], v[1], v[3])
	}
}
