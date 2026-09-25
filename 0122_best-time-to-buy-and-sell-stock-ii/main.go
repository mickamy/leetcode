package main

func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		buy, sell := prices[i-1], prices[i]
		if sell-buy > 0 {
			profit += sell - buy
		}
	}
	return profit
}

//func main() {
//	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
//}
