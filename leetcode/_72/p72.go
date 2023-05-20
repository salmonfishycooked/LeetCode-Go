package _72

const INF = 0x3f3f3f3f

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	f := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		f[i] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		f[i][0] = i
	}
	for i := 0; i <= n2; i++ {
		f[0][i] = i
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
				continue
			}
			f[i][j] = min(min(f[i][j-1], f[i-1][j]), f[i-1][j-1]) + 1
		}
	}
	return f[n1][n2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
