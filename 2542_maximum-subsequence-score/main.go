package main

import (
	"cmp"
	"fmt"
	"slices"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	if n != len(nums2) {
		panic(fmt.Errorf("nums1 length is %d but nums2 has %d length", len(nums1), len(nums2)))
	}

	pairs := make([]pair, n)
	for i := range n {
		pairs[i] = pair{nums1[i], nums2[i]}
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		return cmp.Compare(b.num2, a.num2)
	})

	var sum, ans int
	var heap MinHeap[int]
	for _, p := range pairs {
		heap.Push(p.num1)
		sum += p.num1

		if heap.Len() > k {
			val, _ := heap.Pop()
			sum -= val
		}

		if heap.Len() == k {
			ans = max(ans, sum*p.num2)
		}
	}

	return int64(ans)
}

type pair struct {
	num1, num2 int
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

func main() {
	fmt.Println(maxScore([]int{
		1, 3, 3, 2,
	}, []int{
		2, 1, 3, 4,
	}, 3))
}
