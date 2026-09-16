package main

func moveZeroes(nums []int) {
	lastNonZero := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}
