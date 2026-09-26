package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for k := 0; k <= len(haystack)-len(needle); k++ {
		if needle == haystack[k:k+len(needle)] {
			return k
		}
	}
	return -1
}

//func main() {
//	fmt.Println(strStr("sadbutsad", "sad"))
//	fmt.Println(strStr("leetcode", "leeto"))
//	fmt.Println(strStr("hello", "ll"))
//	fmt.Println(strStr("a", "a"))
//	fmt.Println(strStr("abc", "c"))
//}
