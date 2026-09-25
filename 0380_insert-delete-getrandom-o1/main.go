package main

import (
	"math/rand/v2"
)

type RandomizedSet struct {
	itemMap map[int]int
	items   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{itemMap: make(map[int]int)}
}

func (s *RandomizedSet) Insert(val int) bool {
	_, ok := s.itemMap[val]
	if ok {
		return false
	}
	s.items = append(s.items, val)
	s.itemMap[val] = len(s.items) - 1
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	index, ok := s.itemMap[val]
	if !ok {
		return false
	}

	lastItem := s.items[len(s.items)-1]
	s.items[index] = lastItem
	s.itemMap[lastItem] = index

	s.items = s.items[:len(s.items)-1]
	delete(s.itemMap, val)

	return true
}

func (s *RandomizedSet) GetRandom() int {
	n := rand.IntN(len(s.items))
	return s.items[n]
}
