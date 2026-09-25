package main

import (
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	var profit int
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}
	return profit
}

//func main() {
//	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
//	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
//}
