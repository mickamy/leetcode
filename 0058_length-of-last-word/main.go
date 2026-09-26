package main

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	var last string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		last = word
	}
	return len(last)
}
