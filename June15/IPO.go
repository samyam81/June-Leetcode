package main

import (
	"container/heap"
	"sort"
)

// Project struct to hold capital and profit information
type Project struct {
	capital int
	profit  int
}

// A MaxHeap to hold the profits (we implement this as a min-heap with inverted profits)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Reverses the order to create a max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Project, n)

	for i := 0; i < n; i++ {
		projects[i] = Project{capital: capital[i], profit: profits[i]}
	}

	// Sorting projects based on capital required
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	i := 0

	for j := 0; j < k; j++ {
		// Push all affordable projects into the heap
		for i < n && projects[i].capital <= w {
			heap.Push(maxHeap, projects[i].profit)
			i++
		}

		if maxHeap.Len() == 0 {
			break
		}

		// Get the project with the maximum profit
		w += heap.Pop(maxHeap).(int)
	}

	return w
}