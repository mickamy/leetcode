package main

import (
	"slices"
)

func isIsomorphic(s string, t string) bool {
	occSMap := make(map[string][]int)
	for i, char := range s {
		occSMap[string(char)] = append(occSMap[string(char)], i)
	}

	occTMap := make(map[string][]int)
	for i, char := range t {
		occTMap[string(char)] = append(occTMap[string(char)], i)
	}

	occS := make([][]int, 0, len(occSMap))
	for _, v := range occSMap {
		occS = append(occS, v)
	}

	occT := make([][]int, 0, len(occTMap))
	for _, v := range occTMap {
		occT = append(occT, v)
	}

	slices.SortFunc(occS, func(a, b []int) int {
		return slices.Compare(a, b)
	})
	slices.SortFunc(occT, func(a, b []int) int {
		return slices.Compare(a, b)
	})

	return slices.EqualFunc(occS, occT, func(a []int, b []int) bool {
		return slices.Equal(a, b)
	})
}

func fastIsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var lastSeenS [256]int
	var lastSeenT [256]int

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if lastSeenS[charS] != lastSeenT[charT] {
			return false
		}

		lastSeenS[charS] = i + 1
		lastSeenT[charT] = i + 1
	}

	return true
}
