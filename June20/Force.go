package main

import "sort"

func maxDistance(position []int, m int) int {
    sort.Ints(position)
    low := 1
    high := (position[len(position)-1] - position[0]) / (m - 1)
    answer := 1
    for low <= high {
        mid := low + (high - low) / 2
        if CanWePlace(position, mid, m) {
            answer = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return answer
}

func CanWePlace(arr []int, dist, cow int) bool {
    cntCows := 1
    last := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] - last >= dist {
            cntCows++
            last = arr[i]
        }
        if cntCows >= cow {
            return true
        }
    }
    return false
}