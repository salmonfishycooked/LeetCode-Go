> [LeetCode 199 Binary Tree Right Side View](https://leetcode.cn/problems/binary-tree-right-side-view)

# 解法一：DFS

优先遍历右子树，同时更新答案，如果当前的遍历深度等于答案数组的长度才需要加入答案，否则在当前深度已经有一个节点被先看到了，就不需要加入答案了。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
        f(node.Right, depth + 1)
        f(node.Left, depth + 1)
    }
    f(root, 0)
    return ans
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(n)$



# 解法二：BFS

每次加入当前层的最右节点即可。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    ans := []int{}
    if root == nil {
        return ans
    }

    q := []*TreeNode{root}
    for len(q) > 0 {
        ans = append(ans, q[len(q) - 1].Val)
        help := q
        q = nil
        for _, v := range help {
            if v.Left != nil {
                q = append(q, v.Left)
            }
            if v.Right != nil {
                q = append(q, v.Right)
            }
        }
    }
    return ans
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(n)$