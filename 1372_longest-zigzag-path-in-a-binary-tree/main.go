package main

import "fmt"

func longestZigZag(root *TreeNode) int {
	var maxLen int
	dfs(root, directionLeft, 0, &maxLen)
	dfs(root, directionRight, 0, &maxLen)
	return maxLen
}

type direction string

const (
	directionLeft  direction = "LEFT"
	directionRight direction = "RIGHT"
)

func (d direction) flip() direction {
	switch d {
	case directionLeft:
		return directionRight
	case directionRight:
		return directionLeft
	}
	panic(fmt.Errorf("invalid direction: %s", d))
}

func (d direction) childOf(node TreeNode) *TreeNode {
	switch d {
	case directionLeft:
		return node.Left
	case directionRight:
		return node.Right
	}
	panic(fmt.Errorf("invalid direction: %s", d))
}

func (d direction) String() string {
	return string(d)
}

func dfs(node *TreeNode, d direction, length int, maxLen *int) {
	if node == nil {
		return
	}

	*maxLen = max(*maxLen, length)

	dfs(d.childOf(*node), d.flip(), length+1, maxLen)
	dfs(d.flip().childOf(*node), d, 1, maxLen)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
