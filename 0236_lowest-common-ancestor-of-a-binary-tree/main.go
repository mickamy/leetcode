package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(node *TreeNode, p, q *TreeNode) *TreeNode {
	if node == nil || node == p || node == q {
		return node
	}

	left := dfs(node.Left, p, q)
	right := dfs(node.Right, p, q)

	if left != nil && right != nil {
		return node
	}

	if left != nil {
		return left
	}
	return right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
