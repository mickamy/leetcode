package main

import (
	"slices"
	"strconv"
	"strings"
)

func isHappy(n int) bool {
	var pasts []int
	for {
		digits := numberToDigits(n)
		var sum int
		for i, digit := range digits {
			digits[i] = digit * digit
			sum += digits[i]
		}
		if sum == 1 {
			return true
		}
		if slices.Contains(pasts, sum) {
			return false
		}
		n = sum
		pasts = append(pasts, n)
	}
}

func numberToDigits(n int) []int {
	str := strings.Split(strconv.Itoa(n), "")
	digits := make([]int, len(str))
	for i, s := range str {
		digits[i], _ = strconv.Atoi(s)
	}
	return digits
}

//func main() {
//	fmt.Println(isHappy(19))
//	fmt.Println(isHappy(2))
//}
