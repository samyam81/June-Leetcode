package main

import "sort"
func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize!=0 {
		return false
	}

	count := make(map[int]int)
    for _, card := range hand {
        count[card]++
    }

	sort.Ints(hand)

	for _, card := range hand {
        if count[card] > 0 {
            for i := 0; i < groupSize; i++ {
                if count[card+i] == 0 {
                    return false 
                }
                count[card+i]-- 
            }
        }
    }


	return true
}