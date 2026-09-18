package main

import (
	"slices"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var nodes []*ListNode
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	slices.Reverse(nodes)
	var next *ListNode
	for i := range slices.Backward(nodes) {
		nodes[i].Next = next
		next = nodes[i]
	}

	return nodes[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	vals := []int{1, 2, 3, 4, 5}
	var head, prev *ListNode
	for i, val := range vals {
		newNode := new(ListNode{Val: val})
		if i == 0 {
			head = newNode
		}
		if prev != nil {
			prev.Next = newNode
		}
		prev = newNode
	}
	reverseList(head)
}
