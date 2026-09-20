package main

func minEatingSpeed(piles []int, h int) int {
	low, high := 1, 0
	for _, p := range piles {
		high = max(high, p)
	}

	for low < high {
		mid := low + (high-low)/2
		if canEatAll(piles, h, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func canEatAll(piles []int, h int, k int) bool {
	var hours int
	for _, p := range piles {
		hours += (p + k - 1) / k
	}
	return hours <= h
}
