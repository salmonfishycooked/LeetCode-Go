> [LeetCode 63][https://leetcode.cn/problems/unique-paths]

# 解法：动态规划

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// dp[f(i, j)] 表示机器人走到位置(i, j)的方法数
	// 如果(i, j) 为石头，则令dp[f(i, j)] = 0，表示走到石头这里的方法数为0，经推算这一步是合理的
	dp := make([]int, m*n)

	f := func(i, j int) int {
		return i*n + j
	}

	// 遍历网格第 1 列
	for i := 0; i < m; i++ {
		// 如果遍历到石头，下面所有路都走不通（因为机器人只能向下或向右走）
		if obstacleGrid[i][0] == 1 {
			for ; i < m; i++ {
				dp[f(i, 0)] = 0
			}
			break
		}
		dp[f(i, 0)] = 1
	}
	// 遍历网格第 1 行
	for i := 0; i < n; i++ {
		// 如果遍历到石头，右边所有路都走不通
		if obstacleGrid[0][i] == 1 {
			for ; i < n; i++ {
				dp[f(0, i)] = 0
			}
			break
		}
		dp[f(0, i)] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[f(i, j)] = 0
				continue
			}
			dp[f(i, j)] = dp[f(i-1, j)] + dp[f(i, j-1)]
		}
	}
	return dp[f(m-1, n-1)]
}

```

