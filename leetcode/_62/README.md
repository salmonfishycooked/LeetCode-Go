> [LeetCode 62][https://leetcode.cn/problems/unique-paths]

# 解法：动态规划

```go
func uniquePaths(m int, n int) int {
    // dp[f(i, j)] 表示机器人走到位置(i, j)的方法数
    dp := make([]int, m * n)

    f := func(i, j int) int {
        return i * n + j
    }

    for i := 0; i < m; i++ {
        dp[f(i, 0)] = 1
    }
    for j := 0; j < n; j++ {
        dp[f(0, j)] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[f(i, j)] = dp[f(i - 1, j)] + dp[f(i, j - 1)]
        }
    }
    return dp[f(m - 1, n - 1)]
}
```

