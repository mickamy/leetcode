package main

import (
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func nearestExit(maze [][]byte, entrance []int) int {
	cs := make([][]cell, len(maze))
	var m, n int
	for y, bytes := range maze {
		cs[y] = make([]cell, len(bytes))
		for x, b := range bytes {
			cs[y][x] = cell{
				x:    x,
				y:    y,
				wall: b == '+',
			}
			n = x + 1
		}
		m = y + 1
	}

	ent := position{x: entrance[1], y: entrance[0]}
	return bfs(cells{cells: cs, m: m, n: n}, ent)
}

func bfs(maze cells, entrance position) int {
	queue := Queue[item]{items: []item{{pos: entrance}}}

	visited := make(map[position]bool)
	visited[entrance] = true

	for queue.Len() > 0 {
		curr, _ := queue.Dequeue()

		if curr.pos != entrance && maze.canExit(curr.pos) {
			return curr.step
		}

		for _, next := range []position{curr.pos.left(), curr.pos.right(), curr.pos.up(), curr.pos.down()} {
			if maze.canStepIn(next) && !visited[next] {
				visited[next] = true
				queue.Enqueue(item{pos: next, step: curr.step + 1})
			}
		}
	}

	return -1
}

type position struct {
	x, y int
}

func (p position) left() position {
	return position{x: p.x - 1, y: p.y}
}

func (p position) right() position {
	return position{x: p.x + 1, y: p.y}
}

func (p position) up() position {
	return position{x: p.x, y: p.y - 1}
}

func (p position) down() position {
	return position{x: p.x, y: p.y + 1}
}

type item struct {
	pos  position
	step int
}

type cell struct {
	x, y int
	wall bool
}

type cells struct {
	cells [][]cell
	m, n  int
}

func (cs cells) canExit(p position) bool {
	if !cs.canStepIn(p) {
		return false
	}

	c := cs.cells[p.y][p.x]
	isTop, isBottom := c.y == 0, c.y == cs.m-1
	isLeft, isRight := c.x == 0, c.x == cs.n-1

	return isTop || isBottom || isLeft || isRight
}

func (cs cells) canStepIn(to position) bool {
	if to.x < 0 || to.y < 0 {
		return false
	}
	if to.x > cs.n-1 || to.y > cs.m-1 {
		return false
	}
	return !cs.cells[to.y][to.x].wall
}

func (cs cells) String() {
	cs.Display(nil)
}

func (cs cells) Display(player *position) {
	for y := 0; y < cs.m; y++ {
		for x := 0; x < cs.n; x++ {
			if player != nil && player.x == x && player.y == y {
				fmt.Print("P ")
				continue
			}

			c := cs.cells[y][x]
			if c.wall {
				fmt.Print("█ ")
			} else if cs.canExit(position{x: x, y: y}) {
				fmt.Print("E ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

//func main() {
//	fmt.Println(nearestExit([][]byte{
//		{'+', '.', '+', '+', '+', '+', '+'}, {'+', '.', '+', '.', '.', '.', '+'}, {'+', '.', '+', '.', '+', '.', '+'}, {'+', '.', '.', '.', '.', '.', '+'}, {'+', '+', '+', '+', '.', '+', '.'},
//	}, []int{0, 1}))
//}
