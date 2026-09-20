package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for k := 0; k <= n; k++ {
		dp[0][k] = k
	}

	for i, a := range word1 {
		for k, b := range word2 {
			if a == b {
				dp[i+1][k+1] = dp[i][k]
			} else {
				dp[i+1][k+1] = min(dp[i+1][k], dp[i][k+1], dp[i][k]) + 1
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("horse", "ros")) // 出力: 3
}