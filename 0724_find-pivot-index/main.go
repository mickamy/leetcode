package main

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i, num := range nums {
		rightSum := totalSum - leftSum - num

		if leftSum == rightSum {
			return i
		}

		leftSum += num
	}

	return -1
}
