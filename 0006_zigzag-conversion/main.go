package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	chars := make([][]char, numRows)
	countOfSection := numRows*2 - 2
	var xPos, yPos int
	for i, r := range s {
		chars[yPos] = append(chars[yPos], char{
			value: string(r),
			x:     xPos,
			y:     yPos,
		})
		indexInSection := i % countOfSection
		if indexInSection >= numRows-1 {
			xPos++
			yPos--
		} else {
			yPos++
		}
	}

	var sb strings.Builder
	for y := range chars {
		for x := range chars[y] {
			sb.WriteString(chars[y][x].value)
		}
	}
	return sb.String()
}

type char struct {
	value string
	x, y  int
}

//func main() {
//	fmt.Println(convert("PAYPALISHIRING", 3))
//}
