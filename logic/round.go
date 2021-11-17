package trivia

import (
	"fmt"

	"github.com/willcliffy/ideal-spork/utils"
)

type ROUND string

const (
	RoundOne   = ROUND("Round 1")
	RoundTwo   = ROUND("Round 2")
	RoundThree = ROUND("Round 3")
	Cumulative = ROUND("Cumulative")
	Wagers     = ROUND("Wager")
	RoundFinal = ROUND("Final")
)

type Round struct {
	names  []string
	emails []string
	scores []int
}

func (r *Round) Sort() {
	n := len(r.names)
	if n != len(r.emails) || n != len(r.names) {
		fmt.Printf("Malformed round")
	}

	sortedNames  := make([]string, n)
	sortedEmails := make([]string, n)

	sortedScores, indices := utils.BubbleSortWithIndices(r.scores)

	for i := 0; i < n; i++ {
		sortedNames[i]  = r.names[indices[i]]
		sortedEmails[i] = r.emails[indices[i]]
	}

	r.emails = sortedEmails
	r.names  = sortedNames
	r.scores = sortedScores
}

func (r Round) PrettyPrint() {
	r.Sort()
	for i, name := range r.names {
		fmt.Printf("%2v  %v\n", r.scores[i], name)
	}
	fmt.Println()
}
