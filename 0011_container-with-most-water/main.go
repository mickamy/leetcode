package main

func maxArea(height []int) int {
	ans, l, r := 0, 0, len(height)-1
	for l < r {
		left, right := height[l], height[r]
		x := r - l
		y := min(left, right)
		area := x * y
		ans = max(ans, area)

		if left < right {
			l++
		} else {
			r--
		}
	}
	return ans
}
