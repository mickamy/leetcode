package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var ans int
	for i := range n {
		if !visited[i] {
			ans++
			dfs(i, isConnected, visited)
		}
	}
	return ans
}

func dfs(city int, isConnected [][]int, visited []bool) {
	visited[city] = true

	for neighbor := range isConnected {
		if isConnected[city][neighbor] == 1 && !visited[neighbor] {
			dfs(neighbor, isConnected, visited)
		}
	}
}
