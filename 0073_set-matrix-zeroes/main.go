package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	var firstRowHasZero, firstColHasZero bool

	// find zero in first column
	for y := range m {
		if matrix[y][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// find zero in first row
	for x := range n {
		if matrix[0][x] == 0 {
			firstRowHasZero = true
			break
		}
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if matrix[y][x] == 0 {
				matrix[y][0] = 0
				matrix[0][x] = 0
			}
		}
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if matrix[y][0] == 0 || matrix[0][x] == 0 {
				matrix[y][x] = 0
			}
		}
	}

	if firstRowHasZero {
		for x := range n {
			matrix[0][x] = 0
		}
	}

	if firstColHasZero {
		for y := range m {
			matrix[y][0] = 0
		}
	}
}

//func main() {
//	matrix := [][]int{
//		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
//	}
//	setZeroes(matrix)
//	fmt.Println(matrix)
//}
