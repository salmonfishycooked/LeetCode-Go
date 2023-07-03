package _25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	leftPrev := dummy
	left, right := head, head
	for right != nil {
		left = right
		for i := 0; i < k-1; i++ {
			right = right.Next
			if right == nil {
				break
			}
		}
		if right == nil {
			break
		}
		rightN := right.Next
		// 翻转区间节点
		var prev *ListNode
		for node := left; node != rightN; {
			next := node.Next
			node.Next = prev
			prev = node
			node = next
		}
		left.Next = rightN
		leftPrev.Next = right
		leftPrev = left
		right = rightN
	}
	return dummy.Next
}
