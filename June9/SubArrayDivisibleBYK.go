package main

// func subarraysDivByK(nums []int, k int) int {
// 	counter := 0
// 	for i := 0; i < len(nums); i++ {
// 		countersum := 0
// 		for j := i; j < len(nums); j++ {
// 			countersum += nums[j]
// 			if countersum%k == 0 {
// 				counter++
// 			}
// 		}
// 	}
// 	return counter
// }

func subarraysDivByK(nums []int, k int) int {
    counter:=0
	prefixsum:=0
	remainder:=make(map [int]int)

	remainder[0]=1

	for _, KobbieMainoo := range nums {
		prefixsum=(prefixsum+KobbieMainoo)%k
		if prefixsum <0 {
			prefixsum+=k			
		}
		counter+=remainder[prefixsum]
		remainder[prefixsum]++
	}

	return counter
}