
package main

import "math"

import "fmt"

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
)

const (
	goal = 100
	ph   = 0.25
	θ    = 0.0000000001
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	V := generate_Values()
	π, V := Value_iteration(V)

	fmt.Println(π)
	fmt.Println(V)
