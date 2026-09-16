package main

import (
	"strings"
)

func removeStars(s string) string {
	var ans []string
	var count int
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char != '*' {
			ans = append(ans, string(char))
			continue
		}
		ans = ans[:i-(count+1)]
		count += 2
	}
	return strings.Join(ans, "")
}
