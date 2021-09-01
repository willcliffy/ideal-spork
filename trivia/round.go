package trivia

type ROUND string

const (
	RoundOne   = ROUND("Round 1")
	RoundTwo   = ROUND("Round 2")
	RoundThree = ROUND("Round 3")
	Wagers     = ROUND("Wager")
	RoundFinal = ROUND("Final")
)

type Round struct {
	names  []string
	emails []string
	scores []int
}

func (r *Round) Sort() {

}

func (r Round) PrettyPrint() {

}

func (r Round) GetScore(email, name string) int {
	return 0
}
