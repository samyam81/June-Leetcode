package main

func sortColors(nums []int) {
	const red = 0
	const white = 1
	const blue = 2

	left := 0
	mid := 0
	right := len(nums) - 1

	for mid <= right {
		switch nums[mid] {
		case red:
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		case white:
			mid++
		case blue:
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		}
	}
}
