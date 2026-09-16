package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	if k > len(nums) {
		panic(fmt.Errorf("k is greater than nums length: %d", len(nums)))
	}

	intervals := make([]interval, 0, len(nums)-k)
	var ans float64 = math.MinInt
	for i := 0; i <= len(nums)-k; i++ {
		itv := interval{
			start: i,
			end:   i + k,
		}
		if i == 0 {
			for _, n := range nums[i : i+k] {
				itv.sum += n
			}
			intervals = append(intervals, itv)
		} else {
			last := intervals[len(intervals)-1]
			itv.sum = last.sum - nums[i-1] + nums[i+k-1]
			intervals = append(intervals, itv)
		}
		ans = max(ans, itv.average(k))
	}

	return ans
}

type interval struct {
	start, end, sum int
}

func (i interval) average(k int) float64 {
	return float64(i.sum) / float64(k)
}
