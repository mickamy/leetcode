package main

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	minimum, middle := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num <= minimum {
			minimum = num
		} else if num <= middle {
			middle = num
		} else {
			return true
		}
	}

	return false
}
