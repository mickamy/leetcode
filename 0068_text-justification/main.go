package main

import (
	"fmt"
	"slices"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	lines := []line{
		{words: []string{words[0]}, maxWidth: maxWidth},
	}
	for i := 1; i < len(words); i++ {
		if lines[len(lines)-1].canHold(words[i]) {
			lines[len(lines)-1].append(words[i])
		} else {
			lines = append(lines, line{
				words:    []string{words[i]},
				maxWidth: maxWidth,
			})
		}
	}

	ans := make([]string, len(lines))
	for i, l := range lines {
		ans[i] = l.string(i == len(lines)-1)
		fmt.Printf("[%s]\n", ans[i])
	}
	return ans
}

type line struct {
	words    []string
	maxWidth int
}

func (l *line) clone() line {
	return line{slices.Clone(l.words), l.maxWidth}
}

func (l *line) canHold(word string) bool {
	newLine := l.clone()
	newLine.append(word)
	length := newLine.wordsLen() + newLine.spaceLen()
	return length <= l.maxWidth
}

func (l *line) append(word string) {
	l.words = append(l.words, word)
}

func (l *line) spaceLen() int {
	if len(l.words) == 1 && len(l.words[0]) == l.maxWidth {
		return 0
	}
	return max(l.maxWidth-l.wordsLen(), max(len(l.words)-1, 1))
}

func (l *line) spaceAt(i int) int {
	if len(l.words) > 1 && i == len(l.words)-1 {
		return 0
	}
	total := l.spaceLen()
	spaceCount := max(len(l.words)-1, 1)
	space := total / spaceCount
	remainder := total % spaceCount
	if remainder-1 >= i {
		return space + 1
	}
	return space
}

func (l *line) wordsLen() int {
	var length int
	for _, w := range l.words {
		length += len(w)
	}
	return length
}

func (l *line) string(isLast bool) string {
	if isLast {
		s := strings.Join(l.words, " ")
		s += strings.Repeat(" ", l.maxWidth-(l.wordsLen()+(len(l.words)-1)))
		return s
	}

	var ss []string
	for i, word := range l.words {
		ss = append(ss, word)
		ss = append(ss, strings.Repeat(" ", l.spaceAt(i)))
	}
	s := strings.Join(ss, "")
	fmt.Printf("[%s] %d\n", s, len(s))
	return s
}

//func main() {
//fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
//fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
//fmt.Println(fullJustify([]string{"Listen", "to", "many,", "speak", "to", "a", "few."}, 6))
//}
