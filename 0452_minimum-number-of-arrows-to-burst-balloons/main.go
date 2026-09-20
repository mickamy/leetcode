package main

import (
	"cmp"
	"fmt"
	"slices"
)

func findMinArrowShots(points [][]int) int {
	intervals := make([]interval, len(points))
	var maxX int
	for i, point := range points {
		intervals[i] = interval{point[0], point[1]}
		maxX = max(maxX, point[1])
	}
	slices.SortFunc(intervals, func(a, b interval) int {
		if a.end != b.end {
			return cmp.Compare(a.end, b.end)
		}
		return cmp.Compare(a.start, b.start)
	})

	fmt.Println(intervals)

	arrows := 1
	prevEnd := intervals[0].end

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start > prevEnd {
			arrows++
			prevEnd = intervals[i].end
		}
	}

	return arrows
}

type interval struct {
	start, end int
}
