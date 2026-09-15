package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	dest := make([]int, len(flowerbed))
	copy(dest, flowerbed)
	for i := range dest {
		curr := dest[i]
		if curr == 1 {
			continue
		}
		var prev int
		if i > 0 {
			prev = dest[i-1]
		}
		var next int
		if i < len(dest)-1 {
			next = dest[i+1]
		}

		if prev == 0 && next == 0 {
			n--
			dest[i] = 1
		}
	}
	return n <= 0
}
