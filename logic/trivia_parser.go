package trivia

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/willcliffy/ideal-spork/utils"
	"google.golang.org/api/sheets/v4"
)

func Start(spreadsheets *sheets.SpreadsheetsService, spreadsheetId string) {
	t := NewTriviaHandler(spreadsheets, spreadsheetId, NewScoreKeeper())

	t.RoundOne()
	t.RoundTwo()
	t.RoundThree([]string{"Rice A Roni", "Norman Rockwell", "Billy Mayes", "Rick James", "Cougar"})
	t.CalculateCumulativeScores()
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

	fmt.Printf("Round 1\n")
	t.ScoreKeeper.Rounds[RoundOne] = Round{names, emails, scores}
	t.ScoreKeeper.Rounds[RoundOne].PrettyPrint()
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

	fmt.Printf("Round 2\n")
	t.ScoreKeeper.Rounds[RoundTwo] = Round{names, emails, scores}
	t.ScoreKeeper.Rounds[RoundTwo].PrettyPrint()
}

func (t TriviaHandler) RoundThree(answers []string) {
	if len(answers) != 5 {
		panic("round 3 answers has length other than 5")
	}

	res := t.WaitForAllSubmissions(RoundThree, t.NumPlayers)

	var emails []string 
	var names  []string
	var scores []int

	reader := bufio.NewReader(os.Stdin)

	for _, v := range res {
		emails = append(emails, v[1].(string))
		names  = append(names, v[3].(string))
		
		score   := 0
		correct := 0
		ratio   := 0.75

		for i := 0; i < 5; i++ {
			if len(answers[i]) <= 6 {
				ratio = 0.8
			} else if len(answers[i]) >= 10 {
				ratio = 0.7
			}

			if utils.StringMatch(v[4 + i].(string), answers[i], ratio) {
				score   += 5
				correct += 1
				continue
			}

			fmt.Printf("\tExpected: %s\n\tGot:      %s\nAccept? ", answers[i], v[4 + i].(string))
			input, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			if strings.Contains(input, "y") {
				score   += 5
				correct += 1
				continue
			}
		}

		scores = append(scores, score)
	}

	fmt.Printf("\nRound 3\n")
	t.ScoreKeeper.Rounds[RoundThree] = Round{names, emails, scores}
	t.ScoreKeeper.Rounds[RoundThree].PrettyPrint()
}

func (t TriviaHandler) CalculateCumulativeScores() {
	var emails []string 
	var names  []string
	var scores []int

	for i, email := range t.ScoreKeeper.Rounds[RoundThree].emails {
		emails = append(emails, email)
		names = append(names, t.ScoreKeeper.Rounds[RoundThree].names[i])

		scores = append(scores, t.ScoreKeeper.ScoreCumulative(email))
	}

	fmt.Printf("\nCumulative\n")
	t.ScoreKeeper.Rounds[Cumulative] = Round{names, emails, scores}
	t.ScoreKeeper.Rounds[Cumulative].PrettyPrint()
}

func (t TriviaHandler) WaitForWagers() {
	res := t.WaitForAllSubmissions(Wagers, t.NumPlayers)

	var emails []string 
	var names  []string
	var scores []int

	for _, v := range res {
		emails = append(emails, v[1].(string))
		names = append(names, v[3].(string))

		scoreStr := strings.Split(v[4].(string), " / ")[0]
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			panic(err)
		}

		scores = append(scores, score)
	}

	fmt.Printf("Wagers\n")
	t.ScoreKeeper.Rounds[Wagers] = Round{names, emails, scores}
	t.ScoreKeeper.Rounds[Wagers].PrettyPrint()
}

func (t TriviaHandler) RoundFinal(answer string) {
	res := t.WaitForAllSubmissions(RoundFinal, t.NumPlayers)

	var emails []string 
	var names  []string
	var scores []int

	reader := bufio.NewReader(os.Stdin)

	for _, v := range res {
		emails = append(emails, v[1].(string))
		names = append(names, v[3].(string))

		correct := utils.StringMatch(v[4].(string), answer, 0.90)

		if !correct {
			fmt.Printf("\tExpected: %s\n\tGot:      %s\nAccept? ", answer, v[4].(string))
			input, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			if strings.Contains(input, "y") {
				correct = true
			}
		}

		scores = append(scores, t.ScoreKeeper.ScoreFinal(v[1].(string), correct))
	}

	fmt.Printf("\nFinal Scores\n")
	t.ScoreKeeper.Rounds[RoundFinal] = Round{names, emails, scores}
	t.ScoreKeeper.Rounds[RoundFinal].PrettyPrint()
}
