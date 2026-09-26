package main

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var ans int
	for i, char := range s {
		val := romanMap[byte(char)]
		if i+1 < len(s) && val < romanMap[s[i+1]] {
			ans -= val
		} else {
			ans += val
		}
	}
	return ans
}

//func main() {
//	fmt.Println(romanToInt("III"))
//}
