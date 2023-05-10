package _96

func numTrees(n int) int {
	// dp[i] 的意思是由i个节点能组成的二叉树的数量
	dp := make([]int, n+1)
	// dp[0] 空节点也算一颗二叉搜索树
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			// 每次将头节点的左子树放置 j 个节点，右子树放置 i - j - 1 个节点
			// 此时左右子树能组成的组合数就是 dp[j] * dp[i - j - 1]
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
