package main

import "math"

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(grumpy)
	curWin := 0
	maxWin := 0
	i := 0

	for j := 0; j < n; j++ {
		curWin += customers[j] * grumpy[j]
		maxWin = int(math.Max(float64(curWin), float64(maxWin)))
		if j>=minutes-1 {
			curWin-=customers[i] * grumpy[i]
			i++
		}
	}
	result:=maxWin
	for i := 0; i < len(grumpy); i++ {
		result+=customers[i] * (1 - grumpy[i])
	}
	return result
}
