package main

const mod = 1e9 + 7

func numTilings(n int) int {
	if n <= 3 {
		switch n {
		case 1:
			return 1
		case 2:
			return 2
		case 3:
			return 5
		}
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % mod
	}
	return dp[n]
}

//func main() {
//	fmt.Println(numTilings(4)) // 11
//}
