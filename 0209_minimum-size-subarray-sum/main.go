package main

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt
	var l, sum int
	for r := range nums {
		sum += nums[r]

		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if ans == math.MaxInt {
		return 0
	}

	return ans
}

//func main() {
//	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
//}
