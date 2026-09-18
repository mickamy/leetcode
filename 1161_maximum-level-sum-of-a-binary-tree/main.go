package main

import (
	"math"
)

type queue[T any] struct {
	items []T
}

func (q *queue[T]) enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *queue[T]) dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

func (q *queue[T]) len() int {
	return len(q.items)
}

func maxLevelSum(root *TreeNode) int {
	q := queue[*TreeNode]{}
	q.enqueue(root)

	maximum, level, maxLevel := math.MinInt, 1, 0
	for q.len() > 0 {
		var sum int
		count := q.len()
		for range count {
			node, _ := q.dequeue()
			sum += node.Val
			if node.Left != nil {
				q.enqueue(node.Left)
			}
			if node.Right != nil {
				q.enqueue(node.Right)
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
