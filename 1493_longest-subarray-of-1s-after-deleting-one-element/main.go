package main

func longestSubarray(nums []int) int {
	var l, zero, ans int
	for r, right := range nums {
		if right == 0 {
			zero++
		}

		for zero > 1 {
			if nums[l] == 0 {
				zero--
			}
			l++
		}

		ans = max(ans, (r-l+1)-1)
	}
	return ans
}
