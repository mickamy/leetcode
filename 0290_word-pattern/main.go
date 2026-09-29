package main

import (
	"maps"
	"slices"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	mappings := make(map[string]string)
	for i, word := range words {
		if v, ok := mappings[word]; ok {
			if v != string(pattern[i]) {
				return false
			}
		}
		mappings[word] = string(pattern[i])
	}

	occurrences := make([]string, 0, len(mappings))
	for v := range maps.Values(mappings) {
		if slices.Contains(occurrences, v) {
			continue
		}
		occurrences = append(occurrences, v)
	}

	return len(occurrences) == len(mappings)
}

func efficientWordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	pToW := make(map[byte]string)
	wToP := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]

		if val, exists := pToW[p]; exists {
			if val != w {
				return false
			}
		} else {
			pToW[p] = w
		}

		if val, exists := wToP[w]; exists {
			if val != p {
				return false
			}
		} else {
			wToP[w] = p
		}
	}

	return true
}

//func main() {
//	fmt.Println(wordPattern("aba", "cat cat cat dog"))
//	fmt.Println(wordPattern("abba", "dog dog dog dog"))
//}
