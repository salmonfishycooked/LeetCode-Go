package _1373

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treeInfo struct {
	isBST  bool
	minVal int
	maxVal int
	sumVal int
}

const INF = 0x3f3f3f3f

var res int

func maxSumBST(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) *treeInfo {
	if root == nil {
		return &treeInfo{true, INF, -INF, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)

	if left.isBST && right.isBST && root.Val > left.maxVal && root.Val < right.minVal {
		sum := left.sumVal + right.sumVal + root.Val
		res = max(res, sum)
		return &treeInfo{true, min(left.minVal, root.Val), max(right.maxVal, root.Val), sum}
	} else {
		// 只要返回false即可，后面三个数值不重要了，后面不会被计算
		return &treeInfo{false, 0, 0, 0}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
