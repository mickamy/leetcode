package main

import (
	"fmt"
	"strings"
)

var romanMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	var sb strings.Builder
	for {
		m := num / 1000
		sb.WriteString(strings.Repeat(romanMap[1000], m))
		num = num - m*1000

		if num/100 == 4 {
			sb.WriteString(romanMap[100])
			sb.WriteString(romanMap[500])
			num = num - 400
		} else if num/100 == 9 {
			sb.WriteString(romanMap[100])
			sb.WriteString(romanMap[1000])
			num = num - 900
		} else {
			d := num / 500
			sb.WriteString(strings.Repeat(romanMap[500], d))
			num = num - d*500
			c := num / 100
			sb.WriteString(strings.Repeat(romanMap[100], c))
			num = num - c*100
		}

		if num/10 == 4 {
			sb.WriteString(romanMap[10])
			sb.WriteString(romanMap[50])
			num = num - 40
		} else if num/10 == 9 {
			sb.WriteString(romanMap[10])
			sb.WriteString(romanMap[100])
			num = num - 90
		} else {
			fmt.Println("before", num)
			l := num / 50
			fmt.Println("num l", num, l)
			sb.WriteString(strings.Repeat(romanMap[50], l))
			num = num - l*50
			x := num / 10
			sb.WriteString(strings.Repeat(romanMap[10], x))
			num = num - x*10
		}

		if num == 4 {
			sb.WriteString(romanMap[1])
			sb.WriteString(romanMap[5])
		} else if num == 9 {
			sb.WriteString(romanMap[1])
			sb.WriteString(romanMap[10])
		} else {
			v := num / 5
			sb.WriteString(strings.Repeat(romanMap[5], v))
			num = num - v*5
			sb.WriteString(strings.Repeat(romanMap[1], num))
		}
		break
	}

	return sb.String()
}

//func main() {
//fmt.Println(intToRoman(3749)) // MMMDCCXLIX
//fmt.Println(intToRoman(58))   // LVIII
//fmt.Println(intToRoman(1994)) // MCMXCIV
//}
