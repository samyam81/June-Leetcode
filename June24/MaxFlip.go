package main

func minKBitFlips(nums []int, k int) int {
    n := len(nums)
    flipCount := 0
    diff := make([]int, n+1) 
    
    flip := 0 
    for i := 0; i < n; i++ {
        flip += diff[i]
        if (nums[i] + flip) % 2 == 0 { 
            if i + k > n {
                return -1 
            }
            flipCount++
            flip++
            diff[i+k]-- 
        }
    }
    
    return flipCount
}
