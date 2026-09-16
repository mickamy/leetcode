package main

func longestOnes(nums []int, k int) int {
	var l, zero, ans int
	for r := range nums {
		right := nums[r]
		if right == 0 {
			zero++
		}

		for zero > k {
			if nums[l] == 0 {
				zero--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
