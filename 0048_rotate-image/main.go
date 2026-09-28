package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// transpose
	for i := range n {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	fmt.Println(matrix)

	// reverse
	for i := range n {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

//func main() {
//	rotate([][]int{
//		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
//	})
//}
