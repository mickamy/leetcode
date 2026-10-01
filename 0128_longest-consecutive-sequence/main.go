package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	var ans int
	for num := range set {
		if _, hasPrev := set[num-1]; !hasPrev {
			currentNum := num
			currentStreak := 1

			for {
				if _, exists := set[currentNum+1]; exists {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}

			ans = max(ans, currentStreak)
		}
	}

	return ans
}

//func main() {
//	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
//	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
//	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
//}
