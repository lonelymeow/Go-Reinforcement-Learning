
package main

import "github.com/Reinforcement-Learning-Golang/sarsa"

/*
   Mountain car problem by Francisco Enrique Cordova Gonzalez
*/

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"strconv"
	"time"
)

//Legal actions
const (
	action_reverse int = iota - 1
	action_zero
	action_forward
)

//position and velocity bounds
const (
	position_min = -1.2
	position_max = 0.5
	velocity_min = -0.07
	velocity_max = 0.07
)

type State struct {
	position     float64
	velocity     float64
	v            *sarsa.ValueFunction
	posScale     float64
	velScale     float64
	max_position float64
	max_velocity float64
	min_position float64
	min_velocity float64
	hash_table   map[string]int
	max_size     int
}

func NewState() State {
	s := State{position: -0.5, velocity: 0}
	s.v = &sarsa.ValueFunction{}
	s.v.New(1, 2048, 8, 0.4/8)
	s.max_position = 0.5
	s.min_position = -1.2
	s.max_velocity = 0.07
	s.min_velocity = -0.07