package _236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	// 如果在左边找到 p 或者 q 的话 leftFlag 就不会为 nil，右边同理
	leftFlag := lowestCommonAncestor(root.Left, p, q)
	rightFlag := lowestCommonAncestor(root.Right, p, q)
	if leftFlag != nil && rightFlag != nil {
		return root
	}
	if leftFlag != nil {
		return leftFlag
	}
	return rightFlag
}
