package main

import (
	"maps"
	"slices"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	consecutive := make(map[int]*int)
	consecutive[nums[0]] = nil
	currMin := nums[0]
	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]
		if prev+1 == curr {
			currMin = min(currMin, prev)
			consecutive[currMin] = &curr
		} else {
			consecutive[curr] = &curr
			currMin = curr
		}
	}

	keys := slices.Collect(maps.Keys(consecutive))
	slices.Sort(keys)

	ans := make([]string, 0, len(consecutive))
	for _, k := range keys {
		v := consecutive[k]
		if v == nil || k == *v {
			ans = append(ans, strconv.Itoa(k))
		} else {
			ans = append(ans, strconv.Itoa(k)+"->"+strconv.Itoa(*v))
		}
	}

	return ans
}

//func main() {
//	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
//	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
//	fmt.Println(summaryRanges([]int{-1}))
//	fmt.Println(summaryRanges([]int{1, 3}))
//	fmt.Println(summaryRanges([]int{0, 1, 3, 4, 5, 6}))
//	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7, 9, 10, 12, 14, 15, 16, 27, 28}))
//}
