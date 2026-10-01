package main

func containsNearbyDuplicate(nums []int, k int) bool {
	lastSeen := make(map[int]int, len(nums))

	for i, num := range nums {
		if prevIdx, exists := lastSeen[num]; exists {
			if i-prevIdx <= k {
				return true
			}
		}
		lastSeen[num] = i
	}

	return false
}

//func main() {
//	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
//}
