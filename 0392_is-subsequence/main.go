package main

func isSubsequence(s string, t string) bool {
	var idx int
	for _, char := range t {
		if idx < len(s) && byte(char) == s[idx] {
			idx++
		}
	}
	return idx == len(s)
}
