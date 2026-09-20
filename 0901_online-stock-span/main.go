package main

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popped, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

type Pair struct {
	Price int
	Span  int
}
type StockSpanner struct {
	stack Stack[Pair]
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (ss *StockSpanner) Next(price int) int {
	currentSpan := 1

	for ss.stack.Len() > 0 {
		top, _ := ss.stack.Peek()

		if top.Price > price {
			break
		}

		popped, _ := ss.stack.Pop()
		currentSpan += popped.Span
	}

	ss.stack.Push(Pair{
		Price: price,
		Span:  currentSpan,
	})

	return currentSpan
}
