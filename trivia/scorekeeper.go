package trivia

import (
	"fmt"
)

type ScoreKeeper struct {
	Players     []Player

	Round1      Round
	Round2      Round
	Round3      Round
	Cumulative  Round
	Wagers      Round
	FinalScores Round
}

func NewScoreKeeper() ScoreKeeper {
	return ScoreKeeper{}
}

func (s ScoreKeeper) PrettyPrintScores(round ROUND) {
	fmt.Printf("\n%v\n", round)

	switch round {
	case RoundOne:
		for i, name := range s.Round1.names {
			fmt.Printf("%2v  %v\n", s.Round1.scores[i], name)
		}
		return
	case RoundTwo:
		for i, name := range s.Round2.names {
			fmt.Printf("%2v  %v\n", s.Round2.scores[i], name)
		}
		return
	case RoundThree:
		for i, name := range s.Round3.names {
			fmt.Printf("%2v  %v\n", s.Round3.scores[i], name)
		}

		fmt.Printf("\nCumulative\n")
		cumulative := make(map[string]int)
		for i, name := range s.Round1.names {
			cumulative[name] += s.Round1.scores[i]
		}
		// for team, score := range s.Scores[RoundTwo] {
		// 	s.Cumulative[team] += score
		// }
		// for team, score := range s.Scores[RoundThree] {
		// 	s.Cumulative[team] += score
		// }

		// //s.Cumulative.

		// for team, score := range s.Cumulative {
		// 	fmt.Printf("%2v  %v\n", score, team)
		// }
		return
	case Wagers:
		s.Wagers.Sort()
		for i, name := range s.Wagers.names {
			fmt.Printf("%2v  %v\n", s.Wagers.scores[i], name)
		}
		return
	case RoundFinal:
		s.FinalScores.Sort()
		for i, name := range s.FinalScores.names {
			fmt.Printf("%2v  %v\n", s.FinalScores.scores[i], name)
		}
		return
	}
}

func (s *ScoreKeeper) ScoreFinal(name, email string, correct bool) {
	if correct {
		s.FinalScores.names  = append(s.FinalScores.names, name)
		s.FinalScores.emails = append(s.FinalScores.emails, email)
		s.FinalScores.scores = append(s.FinalScores.scores, 
			s.Cumulative.GetScore(name, email) + s.Wagers.GetScore(name, email))
	} else {
		s.FinalScores.names  = append(s.FinalScores.names, name)
		s.FinalScores.emails = append(s.FinalScores.emails, email)
		s.FinalScores.scores = append(s.FinalScores.scores, 
			s.Cumulative.GetScore(name, email) - s.Wagers.GetScore(name, email))
	}
}
