
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Up int = iota
	Down
	Left
	Right
)

type QLearningTD struct {
	Q [][]float64

	Qn int
	Qm int

	Sn int
	Sm int

	ter_n int
	ter_m int

	ini_n int
	ini_m int

	α float64
	ε float64
	γ float64
}

func (q *QLearningTD) Initialize() {

	q.α = 0.5
	q.ε = 0.1
	q.γ = 1

	q.Sn = 4
	q.Sm = 12

	Actions := 4 // up, down, left, right
