
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	max_car_location_1 int = 5
	max_car_location_2     = 5

	λ_request_location_1  = 3
	λ_request_location_2  = 4
	λ_drop_off_location_1 = 3
	λ_drop_off_location_2 = 2

	γ = 0.9
	θ = 0.001

	max_cars_overflow_parking_1 = 10
	min_cars_location_1         = 0
	min_cars_location_2         = 0
	max_cars_overflow_parking_2 = 10
	//max_transferred_cars      = 5
	reward_overflow_parking_1 = -4
	reward_overflow_parking_2 = -4
	reward_rented_car         = 10
	reward_bad_move           = -1000
	reward_transferred_car    = -2

	employee_near_location_2 = true
	money_saved_by_employee  = 2
)

type State struct {
	V float64
	π Action
}

type Action struct {
	action1 int
	action2 int
}

type Mat [max_car_location_1 + 1][max_car_location_2 + 1]State

func main() {
	rand.Seed(time.Now().Unix())
	S := get_all_states()
	S = policy_iteration(S)
	print_mat(S)
	print_cars(S)
}

func print_cars(S Mat) {
	for i := 0; i <= max_car_location_1; i++ {
		for j := 0; j <= max_car_location_2; j++ {
			fmt.Print(math.Abs(float64(S[i][j].π.action1)), " ")
		}
		fmt.Println("")
	}
}

func print_mat(S Mat) {
	for i := 0; i <= max_car_location_1; i++ {
		for j := 0; j <= max_car_location_2; j++ {
			fmt.Print("(i: ", i, " j:", j, " : ", S[i][j])
		}
		fmt.Println("")
	}
}

func update_π(S Mat) (bool, Mat) {
	stable := true
	for i := 0; i <= max_car_location_1; i++ {
		for j := 0; j <= max_car_location_2; j++ {
			_action := S[i][j].π

			_i := int(math.Min(float64(i+_action.action1), float64(max_car_location_1)))
			_j := int(math.Min(float64(j+_action.action2), float64(max_car_location_2)))
			max := S[_i][_j].V
			for _, a := range get_actions(i, j) {