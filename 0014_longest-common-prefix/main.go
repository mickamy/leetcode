package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	common := strs[0]
	for i := 1; i < len(strs); i++ {
		s := strs[i]
		for k := 0; k < len(common); k++ {
			if k < len(s) && common[k] == s[k] {
				continue
			}
			common = common[:k]
			break
		}
	}

	return common
}

//func main() {
//	fmt.Println(longestCommonPrefix([]string{
//		"flower", "flow", "flight",
//	}))
//}
