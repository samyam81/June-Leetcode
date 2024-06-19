package main

func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }
    
    left, right := 1, 1_000_000_000
    result := -1
    
    for left <= right {
        mid := left + (right - left) / 2
        if canMakeBouquets(bloomDay, m, k, mid) {
            result = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    return result
}

func canMakeBouquets(bloomDay []int, m int, k int, days int) bool {
    flowers := 0
    bouquets := 0
    
    for _, bloom := range bloomDay {
        if bloom <= days {
            flowers++
            if flowers == k {
                bouquets++
                flowers = 0
                if bouquets == m {
                    return true
                }
            }
        } else {
            flowers = 0
        }
    }
    
    return false
}
