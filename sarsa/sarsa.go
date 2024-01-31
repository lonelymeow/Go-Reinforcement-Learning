package sarsa

type State interface {
	GetRandomFirstPosition() State
	GetActions() []strin