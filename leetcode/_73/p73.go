package _73

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	colZero, rowZero := false, false
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	if rowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
