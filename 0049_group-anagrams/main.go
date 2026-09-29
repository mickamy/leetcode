package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		anagrams[sorted] = append(anagrams[sorted], s)
	}

	ans := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		ans = append(ans, v)
	}

	return ans
}

func sortString(s string) string {
	b := []byte(s)
	slices.Sort(b)
	return string(b)
}

//func main() {
//	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
//	fmt.Println(groupAnagrams([]string{"c", "c"}))
//}
