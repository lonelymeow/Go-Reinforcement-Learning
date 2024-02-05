package sarsa

type State interface {
	GetRandomFirstPosition() State
	GetActions() []string
	GetActiveTiles(string) [][]int
	InGoal