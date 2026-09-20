package main

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i, a := range text1 {
		for k, b := range text2 {
			if a == b {
				dp[i+1][k+1] = dp[i][k] + 1
			} else {
				dp[i+1][k+1] = max(dp[i][k+1], dp[i+1][k])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

//func main() {
//	fmt.Println(longestCommonSubsequence("abcde", "ace"))
//}
