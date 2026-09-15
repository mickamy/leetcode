package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)

	for i := len2; i > 0; i-- {
		candidate := str2[:i]
		if strings.Repeat(candidate, len1/i) == str1 &&
			strings.Repeat(candidate, len2/i) == str2 {
			return candidate
		}
	}

	return ""
}
