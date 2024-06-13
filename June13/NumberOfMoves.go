package main

import (
	"sort" 
	"math"
)


func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	var totalMoves int

	if len(seats) != len(students) {
		return -1 
	}

	for i := 0; i < len(seats); i++ {
		totalMoves += int(math.Abs(float64(seats[i] - students[i])))
	}

	return totalMoves
}