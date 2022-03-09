
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