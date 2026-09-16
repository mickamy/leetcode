package main

func maxOperations(nums []int, k int) int {
	occurrences := make(map[int]int, len(nums))
	for _, num := range nums {
		occurrences[num]++
	}

	var ans int
	for a := range occurrences {
		b := k - a
		if a == b {
			for occurrences[a] >= 2 {
				occurrences[a] -= 2
				ans++
			}
			continue
		}
		for occurrences[a] > 0 && occurrences[b] > 0 {
			occurrences[a]--
			occurrences[b]--
			ans++
		}
	}

	return ans
}
