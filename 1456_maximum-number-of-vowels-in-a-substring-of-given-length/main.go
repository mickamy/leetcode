package main

func maxVowels(s string, k int) int {
	var ans, count int
	for i := 0; i <= len(s)-k; i++ {
		if i == 0 {
			for _, c := range s[i : i+k] {
				if isVowel(c) {
					count++
				}
			}
		} else {
			if isVowel(rune(s[i-1])) {
				count = max(0, count-1)
			}
			if isVowel(rune(s[i+k-1])) {
				count++
			}
		}
		ans = max(ans, count)
	}

	return ans
}

func isVowel(char rune) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
