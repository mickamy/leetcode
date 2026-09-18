package main

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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	q := queue[*TreeNode]{}
	q.enqueue(root)

	for q.len() > 0 {
		levelSize := q.len()

		for i := range levelSize {
			node, _ := q.dequeue()

			if i == levelSize-1 {
				ans = append(ans, node.Val)
			}

			if node.Left != nil {
				q.enqueue(node.Left)
			}
			if node.Right != nil {
				q.enqueue(node.Right)
			}
		}
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
