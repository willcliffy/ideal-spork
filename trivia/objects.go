package trivia

type ROUND string

const (
	RoundOne   = ROUND("Round 1")
	RoundTwo   = ROUND("Round 2")
	RoundThree = ROUND("Round 3")
	RoundFinal = ROUND("Final")
)

type Player struct {
	TeamName string
	Emails   []string
}
