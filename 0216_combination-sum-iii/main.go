package main

import (
	"fmt"
	"slices"
)

func combinationSum3(k int, n int) [][]int {
	var combs []combinations
	backtrack(k, n, &combs)
	ans := make([][]int, len(combs))
	for i, comb := range combs {
		ans[i] = comb.nums
	}
	return ans
}

func backtrack(k, n int, combs *[]combinations) {
	if len(*combs) == 0 {
		*combs = append(*combs, combinations{make([]int, 0, 3)})
	}

	comb := (*combs)[len(*combs)-1]
	if comb.isFulfilled(k) {
		return
	}

	for i := range 10 {
		if comb.contains(i) {
			continue
		}
		backtrack(k, n-i, combs)
	}
}

type combinations struct {
	nums []int
}

func (c combinations) contains(i int) bool {
	return slices.Contains(c.nums, i)
}

func (c combinations) isFulfilled(k int) bool {
	return len(c.nums) == k
}

func (c combinations) sum() int {
	var sum int
	for _, num := range c.nums {
		sum += num
	}
	return sum
}

func (c combinations) len() int {
	return len(c.nums)
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
