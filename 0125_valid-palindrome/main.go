package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if isAlphaNumeric(s[i]) {
			sb.WriteString(strings.ToLower(string(s[i])))
			continue
		}
	}

	str := sb.String()
	palindrome := true
	for i := 0; i < len(str)/2; i++ {
		left, right := str[i], str[len(str)-i-1]
		if left == right {
			continue
		}
		palindrome = false
		break
	}
	return palindrome
}

func isAlphaNumeric(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

//func main() {
//	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
//	fmt.Println(isPalindrome("0P"))
//}
