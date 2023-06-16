package _199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	var f func(node *TreeNode, depth int)
	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth >= len(ans) {
			ans = append(ans, node.Val)
		}
		f(node.Right, depth+1)
		f(node.Left, depth+1)
	}
	f(root, 0)
	return ans
}
