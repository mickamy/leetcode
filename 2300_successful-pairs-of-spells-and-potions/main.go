package main

import (
	"slices"
)

func successfulPairs(
	spells []int,
	potions []int,
	success int64,
) []int {
	slices.Sort(potions)

	ans := make([]int, len(spells))
	for i, spell := range spells {
		count := searchSuccessCount(potions, spell, int(success))
		ans[i] = count
	}

	return ans
}

func searchSuccessCount(potions []int, spell, success int) int {
	l, h, idx := 0, len(potions)-1, len(potions)

	for l <= h {
		mid := l + (h-l)/2
		if potions[mid]*spell >= success {
			idx = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return len(potions) - idx
}

//func main() {
//	fmt.Println(successfulPairs(
//		[]int{5, 1, 3},
//		[]int{1, 2, 3, 4, 5},
//		7,
//	))
//}
