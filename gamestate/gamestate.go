package gamestate

const (
	Unkown = 0
	Menu   = 1
	P1Turn = 2
	P2Turn = 3
)

var currentState int

func SetState(state int) {
	currentState = state
}

func GetState() int {
	return currentState
}
