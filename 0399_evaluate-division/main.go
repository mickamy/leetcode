package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]edge)
	for i, chars := range equations {
		u, v, w := chars[0], chars[1], values[i]
		graph[u] = append(graph[u], edge{to: v, weight: w})
		graph[v] = append(graph[v], edge{to: u, weight: 1.0 / w})
	}

	ans := make([]float64, len(queries))
	for i, chars := range queries {
		a, b := chars[0], chars[1]
		if _, ok := graph[a]; !ok {
			ans[i] = -1
		} else if _, ok := graph[b]; !ok {
			ans[i] = -1
		} else {
			ans[i] = dfs(a, b, 1, graph, make(map[string]bool))
		}
	}
	return ans
}

func dfs(curr string, target string, acc float64, g map[string][]edge, visited map[string]bool) float64 {
	if curr == target {
		return acc
	}

	visited[curr] = true

	for _, e := range g[curr] {
		if !visited[e.to] {
			res := dfs(e.to, target, acc*e.weight, g, visited)
			if res != -1.0 {
				return res
			}
		}
	}

	return -1.0
}

type edge struct {
	to     string
	weight float64
}
