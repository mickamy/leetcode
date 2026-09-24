package main

func majorityElement(nums []int) int {
	occurrences := make(map[int]int)
	for _, num := range nums {
		occurrences[num]++
	}

	var majority, maximum int
	for k, v := range occurrences {
		if v > maximum {
			maximum = v
			majority = k
		}
	}
	return majority
}

//func main() {
//	fmt.Println(majorityElement([]int{3, 2, 3}))
//}
