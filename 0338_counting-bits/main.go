package main

import (
	"strconv"
)

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i/2] + (i % 2)
	}

	return ans
}

func count1InBinary(n int) int {
	if n == 0 {
		return 0
	}

	var count int
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}

	return count
}

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	var binary string
	for n > 0 {
		remainder := n % 2
		binary = strconv.Itoa(remainder) + binary
		n = n / 2
	}

	return binary
}

//func main() {
//	fmt.Println(countBits(2))
//}
