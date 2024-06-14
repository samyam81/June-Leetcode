package main

import "sort"

func minIncrementForUnique(nums []int) int {
    if len(nums)==0 {
		return -1
	}

	sort.Ints(nums)
	var tracker uint=0
	for i := 1; i < len(nums); i++ {
		if nums[i]<=nums[i-1]{
			tracker=tracker+uint(nums[i-1]-nums[i]+1)
			nums[i]=nums[i-1]+1
		}
	}

	return int(tracker)
}