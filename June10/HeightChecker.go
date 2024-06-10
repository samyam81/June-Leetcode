package main

import "sort"

func heightChecker(heights []int) int {
	n := len(heights)
	var count uint = 0
	expected := make([]int, n)

	for i := 0; i < n; i++ {
		expected[i] = heights[i]
	}

	sort.Ints(expected)

	for i := 0; i < n; i++ {
		if heights[i]!=expected[i] {
			count++		
		}
	}
	return int(count)
}