package main

import (
	"slices"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rows, columns := make([][]int, n), make([][]int, n)
	for i := range columns {
		columns[i] = make([]int, len(grid[0]))
	}
	for y, row := range grid {
		rows[y] = row
		for x, column := range row {
			columns[x][y] = column
		}
	}

	var ans int
	for _, row := range rows {
		for i := range n {
			allSame := true
			if !slices.Equal(row, columns[i]) {
				allSame = false
			}

			if allSame {
				ans++
			}
		}
	}

	return ans
}
