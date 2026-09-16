package main

import "strconv"

func compress(chars []byte) int {
	var write, read int

	for read < len(chars) {
		curr := chars[read]
		var count int

		for read < len(chars) && chars[read] == curr {
			read++
			count++
		}

		chars[write] = curr
		write++

		if count > 1 {
			for _, c := range strconv.Itoa(count) {
				chars[write] = byte(c)
				write++
			}
		}
	}

	return write
}
