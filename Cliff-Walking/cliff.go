
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

	q.ini_n = 0 //Start
	q.ini_m = 0

	q.ter_n = 0 //Goal
	q.ter_m = 11

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

}

func (q *QLearningTD) Pi() {
	for i := q.Sn - 1; i >= 0; i-- {
		for j := 0; j < q.Sm; j++ {
			if i == q.ter_n && j == q.ter_m {
				fmt.Print(" G")
			} else {
				switch q.GetAction(i, j) {
				case Up:
					fmt.Print(" U")
				case Down:
					fmt.Print(" D")
				case Left:
					fmt.Print(" L")
				case Right:
					fmt.Print(" R")