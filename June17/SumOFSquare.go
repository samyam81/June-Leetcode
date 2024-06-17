package main

import("math")

func judgeSquareSum(c int) bool {

	if c<0 {
		return false
	}
	right:=int(math.Sqrt(float64(c)))
	left:=0

	for left<=right{
		sum:=left*left+right*right

		if sum==c {
			return true
		} else if sum<c {
			left++
		} else {
			right--
		}
	}
	return false
}