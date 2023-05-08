package _53

func maxSubArray(nums []int) (ans int) {
	n := len(nums)
	dp := make([]int, n)
	// dp[i] 表示包含 nums[i] 且由 nums[i] 前面的数所组成的子数组中最大的和
	// 状态转移方程 dp[i] = max(nums[i], dp[i - 1] + nums[i])
	dp[0] = nums[0]
	ans = dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
