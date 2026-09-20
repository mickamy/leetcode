package main

func maxProfit(prices []int, fee int) int {
	hold, free := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]

		nextHold := max(hold, free-price)
		nextFree := max(free, hold+price-fee)

		hold = nextHold
		free = nextFree
	}

	return free
}

//func main() {
//	fmt.Println(maxProfit(
//		[]int{1, 3, 2, 8, 4, 9},
//		2,
//	))
//}
