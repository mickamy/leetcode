package main

import "fmt"

func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	var ans int
	first, second := head, prev
	for second != nil {
		ans = max(ans, first.Val+second.Val)
		first = first.Next
		second = second.Next
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	vals := []int{4, 2, 2, 3}
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
	fmt.Println(pairSum(head))
}
