package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			ans = append(ans, matrix[top][c])
		}
		top++

		for r := top; r <= bottom; r++ {
			ans = append(ans, matrix[r][right])
		}
		right--

		if top <= bottom {
			for c := right; c >= left; c-- {
				ans = append(ans, matrix[bottom][c])
			}
			bottom--
		}

		if left <= right {
			for r := bottom; r >= top; r-- {
				ans = append(ans, matrix[r][left])
			}
			left++
		}
	}

	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})) // [1 2 3 6 9 8 7 4 5]
}
