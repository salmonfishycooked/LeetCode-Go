> [LeetCode 25 Reverse Nodes in k-Group](https://leetcode.cn/problems/reverse-nodes-in-k-group/)

# 思路

讲一下我的思路，将 `left` 定在待翻转的区间左边，`right` 定在待翻转的区间右边，同时需要记录一开始 `left` 左边的节点（记作 `leftPrev` ）和 `right` 右边的节点（记作 `rightN` ），将区间 `[left, right]` 翻转，然后将 `leftPrev` 和 `rightN` 连上正确的位置，更新变量值，然后继续翻转下一个区间。

完整代码如下：

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    leftPrev := dummy
    left, right := head, head
    for right != nil {
        left = right
        for i := 0; i < k - 1; i++ {
            right = right.Next
            if right == nil {
                break
            }
        }
        // 区间凑不够 k 个数了，直接 break 循环
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
```

- 时间复杂度：$O(n)$
- 空间复杂度：$O(1)$