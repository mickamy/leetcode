package main

import (
	"slices"
)

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	for i := range n {
		if len(board[i]) != n {
			return false
		}
	}

	sudoku := make([][]string, n)
	for y, bytes := range board {
		sudoku[y] = make([]string, n)
		for x, char := range bytes {
			sudoku[y][x] = string(char)
		}
	}

	for y := range sudoku {
		for x := range sudoku[y] {
			// horizontal
			for i, horizontal := range sudoku[y] {
				if i == x || horizontal == "." {
					continue
				}
				if horizontal == sudoku[y][x] {
					return false
				}
			}

			// vertical
			vertical := make([]string, n)
			for i := range n {
				if i == y || sudoku[i][x] == "." {
					continue
				}
				vertical[i] = sudoku[i][x]
			}
			if slices.Contains(vertical, sudoku[y][x]) {
				return false
			}

			// sub-box
			box := subBox(sudoku, x, y)
			if slices.Contains(box, sudoku[y][x]) {
				return false
			}
		}
	}

	return true
}

func subBox(sudoku [][]string, x, y int) []string {
	startX, startY := x-x%3, y-y%3
	box := make([]string, 0, x)
	for boxY := startY; boxY < startY+3; boxY++ {
		for boxX := startX; boxX < startX+3; boxX++ {
			if x == boxX && y == boxY || sudoku[boxY][boxX] == "." {
				continue
			}
			box = append(box, sudoku[boxY][boxX])
		}
	}
	return box
}

//func main() {
//	fmt.Println(
//		isValidSudoku(
//			[][]byte{
//				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
//				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
//				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
//				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
//				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
//				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
//				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
//				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
//				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
//			},
//		),
//	)
//}
