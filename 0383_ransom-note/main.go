package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	occurrences := make(map[string]int)
	for _, char := range magazine {
		occurrences[string(char)]++
	}

	for _, char := range ransomNote {
		occurrences[string(char)]--
	}

	fmt.Println(occurrences)

	able := true
	for _, v := range occurrences {
		if v < 0 {
			able = false
			break
		}
	}

	return able
}
