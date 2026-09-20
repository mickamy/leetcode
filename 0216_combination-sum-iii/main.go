package main

import (
	"slices"
)

func combinationSum3(k int, n int) [][]int {
	var combs []combinations
	backtrack(k, n, combinations{make([]int, 0, 3)}, &combs)
	ans := make([][]int, len(combs))
	for i, comb := range combs {
		ans[i] = comb.nums
	}
	return ans
}

func backtrack(k, n int, curr combinations, combs *[]combinations) {
	if curr.isFulfilled(k, n) {
		*combs = append(*combs, curr)
		return
	}
	if curr.len() == k {
		return
	}

	for i := curr.last() + 1; i < 10; i++ {
		if curr.contains(i) {
			continue
		}
		if curr.sum()+i > n {
			break
		}
		backtrack(k, n, curr.appended(i), combs)
	}
}

type combinations struct {
	nums []int
}

func (c combinations) contains(i int) bool {
	return slices.Contains(c.nums, i)
}

func (c combinations) isFulfilled(k, n int) bool {
	return len(c.nums) == k && c.sum() == n
}

func (c combinations) sum() int {
	var sum int
	for _, num := range c.nums {
		sum += num
	}
	return sum
}

func (c combinations) last() int {
	if len(c.nums) == 0 {
		return 0
	}
	return c.nums[len(c.nums)-1]
}

func (c combinations) appended(num int) combinations {
	nums := make([]int, len(c.nums))
	copy(nums, c.nums)
	nums = append(nums, num)
	return combinations{nums}
}

func (c combinations) len() int {
	return len(c.nums)
}

//func main() {
//	fmt.Println(combinationSum3(4, 1))
//}
