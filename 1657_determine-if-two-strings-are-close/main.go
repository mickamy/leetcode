package main

import (
	"maps"
	"slices"
)

func closeStrings(word1 string, word2 string) bool {
	occ1, occ2 := make(map[int32]int), make(map[int32]int)
	for _, char := range word1 {
		occ1[char]++
	}
	for _, char := range word2 {
		occ2[char]++
	}

	if len(occ1) != len(occ2) {
		return false
	}

	for k := range occ1 {
		if _, ok := occ2[k]; !ok {
			return false
		}
	}

	occ1Vals := slices.Sorted(maps.Values(occ1))
	occ2Vals := slices.Sorted(maps.Values(occ2))

	return slices.Equal(occ1Vals, occ2Vals)
}
