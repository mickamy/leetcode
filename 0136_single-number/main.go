package main

func singleNumber(nums []int) int {
	var ans int
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

//func main() {
//	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
//}
