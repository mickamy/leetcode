package main

import "slices"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, num := range nums2 {
		nums1[i+m] = num
	}
	slices.Sort(nums1)
}
