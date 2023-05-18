package _5

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	// 初始化第一条对角线
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// 初始化第二条对角线
	ans := string(s[0])
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans = s[i : i+2]
		}
	}

	// 递推（以斜对角线递推）
	for step := 2; step <= n; step++ {
		for i := 0; i < n-step; i++ {
			dp[i][i+step] = dp[i+1][i+step-1] && s[i] == s[i+step]
			// 如果该字串是回文串并且长度大于之前所记录的最长长度，则更新 ans
			if dp[i][i+step] && step+1 > len(ans) {
				ans = s[i : i+step+1]
			}
		}
	}

	return ans
}
