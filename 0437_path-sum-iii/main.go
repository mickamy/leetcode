package main

import (
	"slices"
)

func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, []int{}, targetSum)
}

func dfs(node *TreeNode, route []int, target int) int {
	if node == nil {
		return 0
	}

	route = append(route, node.Val)

	count := goodSumCount(route, target)
	count += dfs(node.Left, route, target)
	count += dfs(node.Right, route, target)
	route = route[:len(route)-1]
	return count
}

func goodSumCount(route []int, target int) int {
	var count, sum int
	for _, val := range slices.Backward(route) {
		sum += val
		if sum == target {
			count++
		}
	}
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
