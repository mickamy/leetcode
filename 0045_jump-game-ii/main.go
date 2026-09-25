package main

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var count, i int
	for i < len(nums) {
		var maximum, next int
		for k := 1; k <= nums[i]; k++ {
			if i+k >= len(nums)-1 {
				next = len(nums)
				break
			}
			if k+nums[i+k] >= maximum {
				maximum = k + nums[i+k]
				next = i + k
			}
		}
		i = next
		count++
	}
	return count
}

//func main() {
//	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
//	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
//	fmt.Println(jump([]int{1, 1, 1, 1}))
//}
