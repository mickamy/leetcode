package main

func asteroidCollision(asteroids []int) []int {
	s := stack[int]{}
	for _, a := range asteroids {
		if a > 0 {
			s.push(a)
			continue
		}
		var destroyed bool
		for s.len() > 0 {
			top, _ := s.peek()
			if top < 0 {
				break
			}
			if -a > top {
				s.pop()
				continue
			}
			if -a == top {
				s.pop()
				destroyed = true
				break
			}
			if -a < top {
				destroyed = true
				break
			}
		}
		if !destroyed {
			s.push(a)
		}
	}

	return s.elem
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
