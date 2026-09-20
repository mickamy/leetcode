package main

import (
	"cmp"
)

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

type MonotonicFinder[T cmp.Ordered] struct {
	items []T
}

func NewMonotonicFinder[T cmp.Ordered](items []T) MonotonicFinder[T] {
	return MonotonicFinder[T]{items: items}
}

func (f *MonotonicFinder[T]) FindNextGreater() []int {
	indices := make([]int, len(f.items))
	for i := range indices {
		indices[i] = -1
	}

	stack := Stack[int]{}

	for i := range len(f.items) {
		for stack.Len() > 0 {
			topIdx, _ := stack.Peek()
			if f.items[i] <= f.items[topIdx] {
				break
			}

			poppedIdx, _ := stack.Pop()
			indices[poppedIdx] = i
		}

		stack.Push(i)
	}

	return indices
}

func dailyTemperatures(temperatures []int) []int {
	f := NewMonotonicFinder(temperatures)
	nextGraters := f.FindNextGreater()

	ans := make([]int, len(nextGraters))
	for i, grater := range nextGraters {
		if grater == -1 {
			ans[i] = 0
			continue
		}
		ans[i] = grater - i
	}

	return ans
}

//func main() {
//	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
//}
