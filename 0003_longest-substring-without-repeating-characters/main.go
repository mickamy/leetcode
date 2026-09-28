package main

func lengthOfLongestSubstring(s string) int {
	var lastSeen [128]int

	var ans, l int

	for r := 0; r < len(s); r++ {
		char := s[r]

		if lastSeen[char] > l {
			l = lastSeen[char]
		}

		ans = max(ans, r-l+1)
		lastSeen[char] = r + 1
	}

	return ans
}

//func main() {
//	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
//	fmt.Println(lengthOfLongestSubstring("bbbb"))
//	fmt.Println(lengthOfLongestSubstring("pwwkew"))
//	fmt.Println(lengthOfLongestSubstring("eea"))
//	fmt.Println(lengthOfLongestSubstring("1R1T7"))
//}
