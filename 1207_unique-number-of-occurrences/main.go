package main

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, num := range arr {
		occurrences[num]++
	}
	uniqueness := make(map[int]struct{})
	for _, v := range occurrences {
		if _, ok := uniqueness[v]; ok {
			return false
		}
		uniqueness[v] = struct{}{}
	}

	return true
}
