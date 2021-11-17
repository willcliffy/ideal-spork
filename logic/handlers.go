package trivia

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/willcliffy/ideal-spork/utils"
	"google.golang.org/api/sheets/v4"
)

type ScoreKeeper struct {
	Players     []Player

	Rounds map[ROUND]Round
}

func NewScoreKeeper() ScoreKeeper {
	return ScoreKeeper{
		Rounds: make(map[ROUND]Round),
	}
}

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

func (s ScoreKeeper) ScoreCumulative(email string) int {
	r1 := utils.StringArrayIndexOf(s.Rounds[RoundOne].emails, email)
	r2 := utils.StringArrayIndexOf(s.Rounds[RoundTwo].emails, email)
	r3 := utils.StringArrayIndexOf(s.Rounds[RoundThree].emails, email)

	if r1 < 0 || r2 < 0 || r3 < 0 {
		panic(fmt.Sprintf("Email not found: %v", email))
	}

	return s.Rounds[RoundOne].scores[r1] + s.Rounds[RoundTwo].scores[r2] + s.Rounds[RoundThree].scores[r3]
}

func (s ScoreKeeper) ScoreFinal(email string, correct bool) int {
	cumulativeInd := utils.StringArrayIndexOf(s.Rounds[Cumulative].emails, email)
	wagerInd := utils.StringArrayIndexOf(s.Rounds[Wagers].emails, email)

	if cumulativeInd < 0 || wagerInd < 0 {
		panic(fmt.Sprintf("Email not found: %v", email))
	}

	cumulative := s.Rounds[Cumulative].scores[cumulativeInd]
	wager := s.Rounds[Wagers].scores[wagerInd]

	if correct {
		return cumulative + wager
	} else {
		return cumulative - wager
	}
}
