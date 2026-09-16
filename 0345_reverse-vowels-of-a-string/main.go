package main

func reverseVowels(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	for left < right {
		for left < right && !isVowel(bytes[left]) {
			left++
		}
		for left < right && !isVowel(bytes[right]) {
			right--
		}

		if left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
