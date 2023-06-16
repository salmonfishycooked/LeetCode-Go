> [LeetCode 143 Reorderd List](https://leetcode.cn/problems/reorder-list)

# 解法：寻找链表中点 + 链表逆序 + 合并链表

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    // 找到中间节点 slow
    slow := head
    for fast := head; fast != nil && fast.Next != nil; {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 反转中间节点之后的链表
    // 反转后的新头节点为 tHead，尾节点为 slow
    node := slow
    var tHead *ListNode
    for node != nil {
        next := node.Next
        node.Next = tHead
        tHead = node
        node = next
    }

    // 重排链表
    for head != tHead && head.Next != tHead {
        tNext, next := tHead.Next, head.Next
        head.Next = tHead
        tHead.Next = next
        tHead = tNext
        head = next
    }
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(1)$
