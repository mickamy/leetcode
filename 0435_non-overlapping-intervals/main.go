package main

import (
	"cmp"
	"slices"
)

type interval struct {
	start, end int
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	items := make([]interval, len(intervals))
	for i, iv := range intervals {
		items[i] = interval{start: iv[0], end: iv[1]}
	}

	slices.SortFunc(items, func(a, b interval) int {
		return cmp.Compare(a.end, b.end)
	})

	var removed int
	prevEnd := items[0].end

	for i := 1; i < len(items); i++ {
		if items[i].start < prevEnd {
			removed++
		} else {
			prevEnd = items[i].end
		}
	}

	return removed
}
