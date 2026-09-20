package main

func rob(nums []int) int {
	var prev1, prev2 int
	for _, num := range nums {
		current := max(prev1, prev2+num)

		prev2 = prev1
		prev1 = current
	}

	return prev1
}

//func main() {
//	fmt.Println(rob([]int{1})) // 1
//	fmt.Println(rob([]int{1, 2, 3, 1})) // 4
//	fmt.Println(rob([]int{2, 7, 9, 3, 1})) // 12
//}
