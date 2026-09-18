package main

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	var count int
	if node.Val >= maxVal {
		count++
		maxVal = node.Val
	}

	count += dfs(node.Left, maxVal)
	count += dfs(node.Right, maxVal)
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
