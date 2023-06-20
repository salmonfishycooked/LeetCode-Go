package _51

import "strings"

func solveNQueens(n int) (ans [][]string) {
	row := make([]int, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)

	var dfs func(i, msk int)
	dfs = func(i, msk int) {
		if i == n {
			set := make([]string, n)
			for i, v := range row {
				set[i] = strings.Repeat(".", v) + "Q" + strings.Repeat(".", n-v-1)
			}
			ans = append(ans, set)
			return
		}
		for j := 0; j < n; j++ {
			if (msk>>j)&1 == 1 && !diag1[i+j] && !diag2[i-j+n-1] {
				row[i] = j
				diag1[i+j], diag2[i-j+n-1] = true, true
				dfs(i+1, msk^(1<<j))
				diag1[i+j], diag2[i-j+n-1] = false, false
			}
		}
	}
	dfs(0, (1<<n)-1)
	return
}
