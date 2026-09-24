package main

import "slices"

func removeElement(nums []int, val int) int {
	nums = slices.DeleteFunc(nums, func(i int) bool {
		return i == val
	})
	return len(nums)
}
