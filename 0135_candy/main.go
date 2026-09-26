package main

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	var total int
	for _, c := range candies {
		total += c
	}
	return total
}

//func main() {
//	fmt.Println(candy([]int{1, 0, 2}))                   // 5
//	fmt.Println(candy([]int{1, 2, 2}))                   // 4
//	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))   // 13
//}
