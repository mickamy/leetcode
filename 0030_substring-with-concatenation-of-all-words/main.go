package main

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount

	if len(s) < totalLen {
		return nil
	}

	counts := make(map[string]int, wordCount)
	for _, w := range words {
		counts[w]++
	}

	var ans []int

	for i := range wordLen {
		l := i
		seen := make(map[string]int, wordCount)
		var matchedWords int

		for r := i; r <= len(s)-wordLen; r += wordLen {
			word := s[r : r+wordLen]

			if targetCount, exists := counts[word]; exists {
				seen[word]++
				matchedWords++

				for seen[word] > targetCount {
					leftWord := s[l : l+wordLen]
					seen[leftWord]--
					matchedWords--
					l += wordLen
				}

				if matchedWords == wordCount {
					ans = append(ans, l)
				}
			} else {
				clear(seen)
				matchedWords = 0
				l = r + wordLen
			}
		}
	}

	return ans
}

//func main() {
//	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"})) // [0 9]
//}
