package main

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	degree := make([]int, n)
	for _, road := range roads {
		degree[road[0]]++
		degree[road[1]]++
	}

	cities := make([]int, n)
	for i := 0; i < n; i++ {
		cities[i] = i
	}
	sort.Slice(cities, func(i, j int) bool {
		return degree[cities[i]] > degree[cities[j]]
	})

	values := make([]int, n)
	currentValue := n
	for _, city := range cities {
		values[city] = currentValue
		currentValue--
	}

	var totalImportance int64
	for _, road := range roads {
		totalImportance += int64(values[road[0]] + values[road[1]])
	}

	return totalImportance
}