package main

import (
	"fmt"
	"slices"
)

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	next := make([][]int, len(board))
	livings := make([]point, 0, m*n)
	for y := range m {
		next[y] = make([]int, n)
		for x := range n {
			if board[y][x] == 1 {
				livings = append(livings, point{x: x, y: y})
			}
		}
	}

	for y := range board {
		for x := range board[y] {
			p := point{x: x, y: y}

			neighbors := [8]point{p.left().up(), p.up(), p.right().up(), p.right(), p.right().down(), p.down(), p.down().left(), p.left()}
			var living int
			for _, neighbor := range neighbors {
				if !neighbor.included(m, n) {
					continue
				}
				if slices.Contains(livings, neighbor) {
					living++
				}
			}

			fmt.Printf("(%d, %d)=%d\n", y, x, board[y][x])

			if board[y][x] == 1 {
				// Any live cell with fewer than two live neighbors dies as if caused by under-population.
				if living < 2 {
					next[y][x] = 0
				} else
				// Any live cell with two or three live neighbors lives on to the next generation.
				if living == 2 || living == 3 {
					next[y][x] = 1
				} else
				// Any live cell with more than three live neighbors dies, as if by over-population.
				if living > 3 {
					next[y][x] = 0
				}
			} else {
				// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
				if living == 3 {
					next[y][x] = 1
				}
			}
		}
	}

	copy(board, next)
}

type point struct {
	x, y int
}

func (p point) left() point {
	return point{x: p.x - 1, y: p.y}
}

func (p point) right() point {
	return point{x: p.x + 1, y: p.y}
}

func (p point) up() point {
	return point{x: p.x, y: p.y - 1}
}

func (p point) down() point {
	return point{x: p.x, y: p.y + 1}
}

func (p point) included(m, n int) bool {
	return p.x >= 0 && p.x < n && p.y >= 0 && p.y < m
}

//func main() {
//	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
//	gameOfLife(board)
//	fmt.Println(board)
//}
