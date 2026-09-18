package main

import (
	"strings"
)

func decodeString(s string) string {
	countStack := stack[int]{}
	strStack := stack[string]{}
	var curStr string
	var k int

	for _, char := range s {
		switch char {
		case '[':
			countStack.push(k)
			strStack.push(curStr)
			k = 0
			curStr = ""
		case ']':
			str, _ := strStack.pop()
			times, _ := countStack.pop()
			curStr = str + strings.Repeat(curStr, times)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			k = k*10 + int(char-'0')
		default:
			curStr += string(char)
		}
	}
	return curStr
}

type stack[T any] struct {
	elem []T
}

func newStack[T any](elem []T) stack[T] {
	copied := make([]T, 0, len(elem))
	copy(copied, elem)
	return stack[T]{elem: copied}
}

func (s *stack[T]) push(elem T) {
	s.elem = append(s.elem, elem)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.elem) == 0 {
		var zero T
		return zero, false
	}
	top := s.elem[len(s.elem)-1]
	s.elem = s.elem[:len(s.elem)-1]
	return top, true
}

func (s *stack[T]) peek() (T, bool) {
	if len(s.elem) == 0 {
		var zero T
		return zero, false
	}
	return s.elem[len(s.elem)-1], true
}

func (s *stack[T]) len() int {
	return len(s.elem)
}
