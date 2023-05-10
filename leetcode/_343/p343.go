package _343

func integerBreak(n int) int {
	// dp[i]代表将整数i拆分成数个正整数的和，但这种拆分方法拆分出来的整数乘积最大
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i-2; j++ {
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
