package sarsa

type State interface {
	GetRandomFirstPosition() State
	GetActions() []string
	GetActiveTiles(string) [][]int
	InGoalState() bool
	TakeAction(string) (State, float64)
}

type ValueFunction struct {
	Weights  []float64
	Tilings  int
	Alpha    float64
	Features int
}

//constructor
func (v *ValueFunction) New(feature, max_size, tiling int, alpha float64) {

	v.Weights = make([]float64, max_size)

	v.Tilings = tiling
	v.Alpha = alpha
	v.Features = feature

}

type ActionFunction func(State, *ValueFunction) string
type Valuefunction func(State, string, *ValueFunction) float64

func SemiGradientSarsa(state State, GetAction ActionFunction, valueFunction *ValueFunction) int {

	currentState := state.GetRandomFirstPosition()

	cur