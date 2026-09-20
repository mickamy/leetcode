package main

import "cmp"

type SmallestInfiniteSet struct {
	currentSmallest int
	addedBack       MinHeap[int]
	addedSet        map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		currentSmallest: 1,
		addedSet:        make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.addedBack.Len() > 0 {
		smallest, _ := s.addedBack.Pop()
		delete(s.addedSet, smallest)
		return smallest
	}

	ans := s.currentSmallest
	s.currentSmallest++
	return ans
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num >= s.currentSmallest || s.addedSet[num] {
		return
	}
	s.addedSet[num] = true
	s.addedBack.Push(num)
}

type MinHeap[T cmp.Ordered] struct {
	items []T
}

func NewMinHeap[T cmp.Ordered](s []T) *MinHeap[T] {
	items := make([]T, len(s))
	copy(items, s)
	h := &MinHeap[T]{items: items}

	for i := parent(len(items) - 1); i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *MinHeap[T]) Len() int {
	return len(h.items)
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap[T]) Peek() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	return h.items[0], true
}

func (h *MinHeap[T]) Push(v T) {
	h.items = append(h.items, v)
	h.up(len(h.items) - 1)
}

func (h *MinHeap[T]) up(i int) {
	for i > 0 && h.items[parent(i)] > h.items[i] {
		p := parent(i)
		h.items[p], h.items[i] = h.items[i], h.items[p]
		i = p
	}
}

func (h *MinHeap[T]) Pop() (T, bool) {
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

func (h *MinHeap[T]) down(i int) {
	for {
		l, r := leftChild(i), rightChild(i)
		smallest := i

		if l < len(h.items) && h.items[l] < h.items[smallest] {
			smallest = l
		}
		if r < len(h.items) && h.items[r] < h.items[smallest] {
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
	var h MinHeap[T]
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
