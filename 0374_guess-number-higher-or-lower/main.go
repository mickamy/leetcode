package main

func guessNumber(n int) int {
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2

		switch guess(mid) {
		case -1:
			high = mid - 1
		case 1:
			low = mid + 1
		case 0:
			return mid
		}
	}

	return -1
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 */
func guess(num int) int {
	return -1
}
