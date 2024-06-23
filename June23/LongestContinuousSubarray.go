package main

func longestSubarray(nums []int, limit int) int {
	var maxq, minq []int
	n := len(nums)
	j := 0
	ans := 0

	for i := 0; i < n; i++ {

		for len(maxq) > 0 && nums[i] > maxq[len(maxq)-1] {
			maxq = maxq[:len(maxq)-1]
		}
		maxq = append(maxq, nums[i])

		for len(minq) > 0 && nums[i] < minq[len(minq)-1] {
			minq = minq[:len(minq)-1]
		}
		minq = append(minq, nums[i])

		for maxq[0] - minq[0] > limit {
			if nums[j] == maxq[0] {
				maxq = maxq[1:]
			}
			if nums[j] == minq[0] {
				minq = minq[1:]
			}
			j++
		}

		ans = max(ans, i - j + 1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
