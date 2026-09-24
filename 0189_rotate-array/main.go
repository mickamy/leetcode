package main

func rotate(nums []int, k int) {
	dest := make([]int, len(nums))
	for i, num := range nums {
		dest[(i+k)%len(nums)] = num
	}
	copy(nums, dest)
}

//func main() {
//	rotate([]int{
//		1, 2, 3, 4, 5, 6, 7,
//	}, 3)
//}
