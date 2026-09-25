package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	position := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		jump := nums[i]
		if i+jump >= position {
			position = i
		}
	}
	return position == 0
}

//func main() {
//	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
//	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
//}
