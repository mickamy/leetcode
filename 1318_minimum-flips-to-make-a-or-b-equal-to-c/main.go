package main

func minFlips(a int, b int, c int) int {
	flips := 0

	for a > 0 || b > 0 || c > 0 {
		bitA := a & 1
		bitB := b & 1
		bitC := c & 1

		if bitC == 1 {
			if bitA == 0 && bitB == 0 {
				flips++
			}
		} else {
			flips += bitA + bitB
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return flips
}

//func main() {
//	fmt.Println(minFlips(2, 6, 5))
//}
