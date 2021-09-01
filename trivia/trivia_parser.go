package trivia

import (
	"strconv"
	"strings"

	"github.com/willcliffy/ideal-spork/utils"
	"google.golang.org/api/sheets/v4"
)

func Start(spreadsheets *sheets.SpreadsheetsService, spreadsheetId string) {
	t := NewTriviaHandler(spreadsheets, spreadsheetId, NewScoreKeeper())

	t.RoundOne()
	t.RoundTwo()
	t.RoundThree([]string{"Spike Lee", "Ben Franklin", "Sonny Bono", "Marvin Gaye", "Sirhan Sirhan"})
	t.WaitForWagers()
	t.RoundFinal("10")
}

func (t TriviaHandler) RoundOne() {
	res := t.WaitForAllSubmissions(RoundOne, t.NumPlayers)

	var emails []string 
	var names  []string
	var scores []int

	for _, v := range res {
		emails = append(emails, v[1].(string))
		names = append(names, v[3].(string))

		scoreStr := strings.Split(v[2].(string), " / ")[0]
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			panic(err)
		}

		scores = append(scores, score)
	}

	t.ScoreKeeper.Round1 = Round{names, emails, scores}
	t.ScoreKeeper.PrettyPrintScores(RoundOne)
}

func (t TriviaHandler) RoundTwo() {
	res := t.WaitForAllSubmissions(RoundTwo, t.NumPlayers)

	var emails []string 
	var names  []string
	var scores []int

	for _, v := range res {
		emails = append(emails, v[1].(string))
		names = append(names, v[3].(string))

		scoreStr := strings.Split(v[2].(string), " / ")[0]
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			panic(err)
		}

		scores = append(scores, score)
	}

	t.ScoreKeeper.Round2 = Round{names, emails, scores}
	t.ScoreKeeper.PrettyPrintScores(RoundTwo)
}

func (t TriviaHandler) RoundThree(answers []string) {
	if len(answers) != 5 {
		panic("round 3 answers has length other than 5")
	}

	res := t.WaitForAllSubmissions(RoundThree, t.NumPlayers)

	var emails []string 
	var names  []string
	var scores []int

	for _, v := range res {
		emails = append(emails, v[1].(string))
		names = append(names, v[3].(string))
		
		score := 0

		for i := 0; i < 5; i++ {
			if utils.StringMatch(v[4 + i].(string), answers[i], 0.90) {
				score += 5
			}

			// todo - manual scoring for round 3
		}

		scores = append(scores, score)
	}

	t.ScoreKeeper.Round3 = Round{names, emails, scores}
	t.ScoreKeeper.PrettyPrintScores(RoundThree)
}

func (t TriviaHandler) WaitForWagers() {
	res := t.WaitForAllSubmissions(Wagers, t.NumPlayers)

	for _, v := range res {
		t.ScoreKeeper.FinalScores.emails = append(t.ScoreKeeper.FinalScores.emails, v[1].(string))
		t.ScoreKeeper.FinalScores.names = append(t.ScoreKeeper.FinalScores.names, v[3].(string))

		wager, err := strconv.Atoi(v[4].(string))
		if err != nil {
			panic(err)
		}
		
		t.ScoreKeeper.FinalScores.scores = append(t.ScoreKeeper.FinalScores.scores, wager)
	}

	t.ScoreKeeper.PrettyPrintScores(Wagers)
}

func (t TriviaHandler) RoundFinal(answer string) {
	res := t.WaitForAllSubmissions(RoundFinal, t.NumPlayers)

	for _, v := range res {
		// todo - manual scoring for final round
		t.ScoreKeeper.ScoreFinal(
			v[2].(string), 
			v[1].(string), 
			utils.StringMatch(v[3].(string), answer, 0.90))
	}

	t.ScoreKeeper.PrettyPrintScores(RoundFinal)
}
