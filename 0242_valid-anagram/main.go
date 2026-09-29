package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	occS := make(map[rune]int)
	for _, char := range s {
		occS[char]++
	}
	occT := make(map[rune]int)
	for _, char := range t {
		occT[char]++
	}

	for k, v := range occS {
		if occT[k] != v {
			return false
		}
	}

	return true
}