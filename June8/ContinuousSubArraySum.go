package main

func checkSubarraySum(nums []int, k int) bool {
    sum := 0
    sumMap := map[int]int{0: -1} 
    for i, num := range nums {
        sum += num
        if k != 0 {
            sum %= k 
        }
        if prevIndex, found := sumMap[sum]; found {
            if i-prevIndex > 1 {
                return true
            }
        } else {
            sumMap[sum] = i
        }
    }
    return false
}
