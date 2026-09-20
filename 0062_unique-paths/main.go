package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := range m {
		dp[i][0] = 1
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			dp[y][x] = dp[y-1][x] + dp[y][x-1]
		}
	}

	return dp[m-1][n-1]
}

//func main() {
//	fmt.Println(uniquePaths(3, 7))
//	fmt.Println(uniquePaths(3, 2))
//}
