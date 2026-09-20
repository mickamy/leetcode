package main

import (
	"math"
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

func maxLevelSum(root *TreeNode) int {
	q := Queue[*TreeNode]{}
	q.Enqueue(root)

	maximum, level, maxLevel := math.MinInt, 1, 0
	for q.Len() > 0 {
		var sum int
		count := q.Len()
		for range count {
			node, _ := q.Dequeue()
			sum += node.Val
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		if sum > maximum {
			maximum = sum
			maxLevel = level
		}

		level++
	}

	return maxLevel
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
