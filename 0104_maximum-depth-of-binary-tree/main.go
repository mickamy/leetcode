package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}

func dfs(node *TreeNode, depth int) int {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return depth
	}
	return max(dfs(node.Left, depth+1), dfs(node.Right, depth+1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
