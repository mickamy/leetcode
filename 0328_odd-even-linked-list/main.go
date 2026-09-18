package main

import (
	"slices"
)

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var odd, even []*ListNode
	cur := head
	var i int
	for cur != nil {
		if i%2 == 0 {
			even = append(even, cur)
		} else {
			odd = append(odd, cur)
		}
		cur = cur.Next
		i++
	}

	var next *ListNode
	for _, o := range slices.Backward(odd) {
		o.Next = next
		next = o
	}
	for _, e := range slices.Backward(even) {
		e.Next = next
		next = e
	}

	return even[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
