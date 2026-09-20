package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}
	dp := make([]int, len(cost)+2)
	for i := 2; i < len(dp)-1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
		dp[i+1] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

//func main() {
//	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
//	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
//}
