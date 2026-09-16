package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	left, right := make(map[int]bool), make(map[int]bool)
	for _, lhs := range nums1 {
		left[lhs] = true
	}
	for _, rhs := range nums2 {
		right[rhs] = true
	}

	ans := make([][]int, 2)
	for v := range left {
		if _, ok := right[v]; ok {
			continue
		}
		ans[0] = append(ans[0], v)
	}
	for v := range right {
		if _, ok := left[v]; ok {
			continue
		}
		ans[1] = append(ans[1], v)
	}

	return ans
}
