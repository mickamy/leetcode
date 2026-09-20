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

func orangesRotting(g [][]int) int {
	cells := make([][]cell, len(g))
	var m, n int
	for y, ints := range g {
		cells[y] = make([]cell, len(ints))
		for x, i := range ints {
			cells[y][x] = cell{pos: position{x: x, y: y}, state: parseState(i)}
			n = x + 1
		}
		m = y + 1
	}

	return bfs(grid{cells: cells, m: m, n: n})
}

func bfs(g grid) int {
	if g.isAllRotten() {
		return 0
	}

	queue := Queue[cell]{}
	for _, row := range g.cells {
		for _, c := range row {
			if c.state == stateRotten {
				queue.Enqueue(c)
			}
		}
	}

	var step int
	for queue.Len() > 0 {
		var rottens []cell
		for queue.Len() > 0 {
			curr, _ := queue.Dequeue()
			rottens = append(rottens, curr)
		}

		for _, rotten := range rottens {
			for _, next := range []position{rotten.pos.left(), rotten.pos.right(), rotten.pos.up(), rotten.pos.down()} {
				if !g.inBounds(next) {
					continue
				}
				if g.cells[next.y][next.x].rot() {
					queue.Enqueue(g.cells[next.y][next.x])
				}
			}
		}

		step++

		if g.isAllRotten() {
			break
		}
	}

	if !g.isAllRotten() {
		return -1
	}

	return step
}

type state string

const (
	stateNone   state = "NONE"
	stateFresh  state = "FRESH"
	stateRotten state = "ROTTEN"
)

func parseState(i int) state {
	switch i {
	case 0:
		return stateNone
	case 1:
		return stateFresh
	case 2:
		return stateRotten
	}
	panic(fmt.Errorf("invalid state: %d", i))
}

type cell struct {
	pos   position
	state state
}

func (o *cell) rot() bool {
	if o.state == stateNone {
		return false
	}
	if o.state == stateRotten {
		return false
	}
	o.state = stateRotten
	return true
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

type grid struct {
	cells [][]cell
	m, n  int
}

func (g grid) isAllRotten() bool {
	for _, row := range g.cells {
		for _, c := range row {
			if c.state == stateFresh {
				return false
			}
		}
	}
	return true
}

func (g grid) inBounds(to position) bool {
	if to.x < 0 || to.y < 0 {
		return false
	}
	if to.x > g.n-1 || to.y > g.m-1 {
		return false
	}
	return true
}

func (g grid) Display(minute int) {
	fmt.Printf("--- Minute: %d ---\n", minute)
	for y := 0; y < g.m; y++ {
		for x := 0; x < g.n; x++ {
			c := g.cells[y][x]
			switch c.state {
			case stateRotten:
				fmt.Print("R ")
			case stateFresh:
				fmt.Print("F ")
			case stateNone:
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
