
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