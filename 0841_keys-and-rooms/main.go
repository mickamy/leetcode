package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popped, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	visited := make(map[int]struct{})
	visited[0] = struct{}{}

	stack := Stack[int]{items: rooms[0]}
	for stack.Len() > 0 {
		key, _ := stack.Pop()
		if _, ok := visited[key]; ok {
			continue
		}
		visited[key] = struct{}{}
		for _, k := range rooms[key] {
			stack.Push(k)
		}
	}

	return len(visited) == len(rooms)
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{
		{6, 7, 8}, {5, 4, 9}, {}, {8}, {4}, {}, {1, 9, 2, 3}, {7}, {6, 5}, {2, 3, 1},
	}))
}
