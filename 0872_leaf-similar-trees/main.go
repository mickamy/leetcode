package main

import (
	"slices"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var lhs, rhs []int
	dfs(root1, &lhs)
	dfs(root2, &rhs)
	return slices.Equal(lhs, rhs)
}

func dfs(node *TreeNode, dest *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*dest = append(*dest, node.Val)
		return
	}

	dfs(node.Left, dest)
	dfs(node.Right, dest)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
