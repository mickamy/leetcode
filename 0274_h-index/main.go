package main

import (
	"cmp"
	"slices"
)

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	var index int
	for i, citation := range citations {
		if i+1 <= citation {
			index = i + 1
		} else {
			break
		}
	}
	return index
}

//func main() {
//	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
//	fmt.Println(hIndex([]int{1, 3, 1}))
//	fmt.Println(hIndex([]int{100}))
//	fmt.Println(hIndex([]int{0}))
//	fmt.Println(hIndex([]int{0, 0, 2}))
//}
