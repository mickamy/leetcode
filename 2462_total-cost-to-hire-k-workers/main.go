package main

import (
	"cmp"
	"fmt"
)

var less = func(a, b worker) bool {
	if a.cost < b.cost {
		return true
	} else if a.cost > b.cost {
		return false
	}
	return a.index < b.index
}

func totalCost(costs []int, k int, candidates int) int64 {
	left, right := NewAnyMinHeap(less), NewAnyMinHeap(less)

	l, r := 0, len(costs)-1

	for l < candidates && l <= r {
		left.Push(worker{l, costs[l]})
		l++
	}
	for r >= len(costs)-candidates && l <= r {
		right.Push(worker{r, costs[r]})
		r--
	}

	var ans int
	for range k {
		lhs, leftOk := left.Peek()
		rhs, rightOk := right.Peek()

		if leftOk && (!rightOk || less(lhs, rhs)) {
			hired, _ := left.Pop()
			ans += hired.cost
			if l <= r {
				left.Push(worker{l, costs[l]})
				l++
			}
		} else {
			hired, _ := right.Pop()
			ans += hired.cost
			if l <= r {
				right.Push(worker{r, costs[r]})
				r--
			}
		}
	}
	return int64(ans)
}

type worker struct {
	index, cost int
}

type AnyMinHeap[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewAnyMinHeap[T any](less func(a, b T) bool) *AnyMinHeap[T] {
	return &AnyMinHeap[T]{
		items: []T{},
		less:  less,
	}
}

func (h *AnyMinHeap[T]) Len() int {
	return len(h.items)
}

func parent(i int) int     { return (i - 1) / 2 }
func leftChild(i int) int  { return 2*i + 1 }
func rightChild(i int) int { return 2*i + 2 }

func (h *AnyMinHeap[T]) Push(v T) {
	h.items = append(h.items, v)
	h.up(len(h.items) - 1)
}

func (h *AnyMinHeap[T]) up(i int) {
	for i > 0 && h.less(h.items[i], h.items[parent(i)]) {
		p := parent(i)
		h.items[p], h.items[i] = h.items[i], h.items[p]
		i = p
	}
}

func (h *AnyMinHeap[T]) Pop() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}

	top := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]
	h.down(0)
	return top, true
}

func (h *AnyMinHeap[T]) Peek() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	return h.items[0], true
}

func (h *AnyMinHeap[T]) down(i int) {
	for {
		l, r := leftChild(i), rightChild(i)
		smallest := i

		if l < len(h.items) && h.less(h.items[l], h.items[smallest]) {
			smallest = l
		}
		if r < len(h.items) && h.less(h.items[r], h.items[smallest]) {
			smallest = r
		}

		if smallest == i {
			return
		}

		h.items[i], h.items[smallest] = h.items[smallest], h.items[i]
		i = smallest
	}
}

func HeapSort[T cmp.Ordered](s []T) []T {
	var h AnyMinHeap[T]
	for _, v := range s {
		h.Push(v)
	}

	sorted := make([]T, 0, len(s))
	for {
		v, ok := h.Pop()
		if !ok {
			break
		}
		sorted = append(sorted, v)
	}
	return sorted
}

func main() {
	fmt.Println(totalCost([]int{
		17, 12, 10, 2, 7, 2, 11, 20, 8,
	}, 3, 4))
}
