> [LeetCode 102][]

# 解法：广度优先搜索

```go
func levelOrder(root *TreeNode) (ans [][]int) {
    if root == nil {
        return
    }

    q := []*TreeNode{root}
    for len(q) > 0 {
        var tmp []int
        // help 用于存下一层结点
        var help []*TreeNode
        // 将该层结点全部取出
        for len(q) > 0 {
            node := q[0]
            q = q[1:]
            tmp = append(tmp, node.Val)

            // 将该结点的左右结点加入队列
            if node.Left != nil {
                help = append(help, node.Left)
            }
            if node.Right != nil {
                help = append(help, node.Right)
            }
        }
        ans = append(ans, tmp)
        // 将q替换为下一层结点
        q = help
    }
    return
}
```

