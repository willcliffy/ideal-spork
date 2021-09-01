package trivia

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

type TriviaHandler struct {
	NumPlayers  int
	Sheets      *sheets.SpreadsheetsService
	SheetID     string

	Scores      map[ROUND]map[string]int
	Cumulative  map[string]int
	Wagers      map[string]int
	FinalScores map[string]int
}

func (t TriviaHandler) AddScoresToScoreboard(round ROUND, roundScores map[string]int) {
	t.Scores[round] = roundScores
	t.PrettyPrintScores(round)
}

func (t TriviaHandler) PrettyPrintScores(round ROUND) {
	fmt.Println(round)
	for team, score := range t.Scores[round] {
		fmt.Printf("%2v  %v\n", score, team)
	}

	if round == RoundThree {
		for team, score := range t.Scores[RoundOne] {
			t.Cumulative[team] += score
		}
		for team, score := range t.Scores[RoundTwo] {
			t.Cumulative[team] += score
		}
		for team, score := range t.Scores[RoundThree] {
			t.Cumulative[team] += score
		}
		for team, score := range t.Cumulative {
			fmt.Printf("%2v  %v\n", score, team)
		}
	}
}
