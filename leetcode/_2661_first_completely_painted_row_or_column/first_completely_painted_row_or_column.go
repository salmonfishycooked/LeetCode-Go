package _2661_first_completely_painted_row_or_column

func firstCompleteIndex(arr []int, mat [][]int) int {
	pos := make(map[int][2]int)
	// 记录mat中每个元素出现的位置
	for i, _ := range mat {
		for j, v := range mat[i] {
			pos[v] = [2]int{i, j}
		}
	}

	m, n := len(mat), len(mat[0])
	// 用 row, col 记录当前行或当前列已经被涂色的单元格的数量
	row := make([]int, m)
	col := make([]int, n)
	for i, v := range arr {
		x, y := pos[v][0], pos[v][1]
		row[x]++
		col[y]++
		// 判断数v所在矩阵的行或列是否被涂满
		if row[x] == n || col[y] == m {
			return i
		}
	}
	return -1
}
