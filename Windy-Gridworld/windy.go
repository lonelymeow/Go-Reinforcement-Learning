
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

type SarsaTD struct {
	Q [][]float64

	Storm []int

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

func (q *SarsaTD) Initialize() {

	q.α = 0.5
	q.ε = 0.1
	q.γ = 1

	q.Sn = 7
	q.Sm = 10

	Actions := 4 //up, down, left, right

	q.ini_n = 3
	q.ini_m = 0

	q.ter_n = 3
	q.ter_m = 7

	q.Qn = Actions
	q.Qm = q.Sn * q.Sm

	q.Q = make([][]float64, Actions)

	for i := 0; i < Actions; i++ {
		q.Q[i] = make([]float64, q.Sn*q.Sm)
	}

	for i := 0; i < Actions; i++ {
		for j := 0; j < q.Sn*q.Sm; j++ {
			q.Q[i][j] = rand.Float64()
		}
	}

	q.SetQAll(q.ter_n, q.ter_m, 0)

	q.Storm = []int{0, 0, 0, 1, 1, 1, 2, 2, 1, 0}

}

func main() {