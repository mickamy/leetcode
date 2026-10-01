package main

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		prev := merged[len(merged)-1]
		curr := intervals[i]

		if prev[1] >= curr[0] {
			prev[1] = max(prev[1], curr[1])
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

//func main() {
//	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))       // [[1,6],[8,10],[15,18]]
//	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}})) // [[1,10]]
//}
