package main

var digitToLetters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var ans []string
	backtrack(digits, 0, "", &ans)
	return ans
}

func backtrack(digits string, index int, path string, ans *[]string) {
	if index == len(digits) {
		*ans = append(*ans, path)
		return
	}
	letters := digitToLetters[digits[index]]

	for i := 0; i < len(letters); i++ {
		backtrack(digits, index+1, path+string(letters[i]), ans)
	}
}
