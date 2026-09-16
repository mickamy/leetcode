package main

func largestAltitude(gain []int) int {
	var curr, highest int
	for _, diff := range gain {
		curr += diff
		highest = max(highest, curr)
	}

	return highest
}
