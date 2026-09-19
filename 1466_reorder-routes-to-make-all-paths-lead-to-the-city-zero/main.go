package main

import "fmt"

type Edge struct {
	to         int
	isOriginal bool
}

func minReorder(n int, connections [][]int) int {
	graph := make([][]Edge, n)
	for _, c := range connections {
		u, v := c[0], c[1]
		graph[u] = append(graph[u], Edge{to: v, isOriginal: true})
		graph[v] = append(graph[v], Edge{to: u})
	}

	visited := make([]bool, n)
	return dfs(0, graph, visited)
}

func dfs(curr int, graph [][]Edge, visited []bool) int {
	visited[curr] = true

	var changes int
	for _, edge := range graph[curr] {
		if !visited[edge.to] {
			if edge.isOriginal {
				changes++
			}
			changes += dfs(edge.to, graph, visited)
		}
	}

	return changes
}

func main() {
	fmt.Println(minReorder(6, [][]int{
		{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5},
	}))
}
