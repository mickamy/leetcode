package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	slices.Reverse(words)
	dest := make([]string, 0, len(words))
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		dest = append(dest, word)
	}
	return strings.Join(dest, " ")
}
