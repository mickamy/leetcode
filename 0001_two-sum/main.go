package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, a := range nums {
		complement := target - a
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[a] = i
	}
	panic(fmt.Errorf("no answer: %d %d", nums, target))
}
