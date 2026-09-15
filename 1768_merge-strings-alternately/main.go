package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	maxLength := max(len(word1), len(word2))
	var sb strings.Builder
	for i := range maxLength {
		if i < len(word1) {
			sb.WriteString(string(word1[i]))
		}
		if i < len(word2) {
			sb.WriteString(string(word2[i]))
		}
	}
	return sb.String()
}
