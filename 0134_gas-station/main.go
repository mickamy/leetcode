package main

func canCompleteCircuit(gas []int, cost []int) int {
	var total, current, start int

	for i := range len(gas) {
		diff := gas[i] - cost[i]
		total += diff
		current += diff

		if current < 0 {
			start = i + 1
			current = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}

func circularIndex(i, n int) int {
	return (i%n + n) % n
}

//func main() {
//	fmt.Println(canCompleteCircuit(
//		[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2},
//		[]int{2, 3, 4}, []int{3, 4, 3},
//		[]int{5, 8, 2, 8}, []int{6, 5, 6, 6},
//	))
//}
