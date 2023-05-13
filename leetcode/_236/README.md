> [LeetCode 236][]

# 解法一：模拟

找出从根节点到目标结点的路径并记录下来。

对比两条路径，任选一条路径的终点开始从后往前遍历，如果在另一条路径中找到正在遍历中的结点，则该结点就是答案。

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var tmp []*TreeNode
    var ways [][]*TreeNode
    var findWay func(root, target *TreeNode)
    // 找到从 root 结点去到 target 结点的路径
    findWay = func(root, target *TreeNode) {
        tmp = append(tmp, root)
        if root == target {
            trip := make([]*TreeNode, len(tmp))
            copy(trip, tmp)
            ways = append(ways, trip)
        }

        if root.Left != nil {
            tmp = append(tmp, root.Left)
            findWay(root.Left, target)
            tmp = tmp[:len(tmp) - 1]
        }
        if root.Right != nil {
            tmp = append(tmp, root.Right)
            findWay(root.Right, target)
            tmp = tmp[:len(tmp) - 1]
        }

        tmp = tmp[:len(tmp) - 1]
    }
    findWay(root, p)
    findWay(root, q)

    // 遍历两条路径
    way1, way2 := ways[0], ways[1]
    for i := len(way1) - 1; i >= 0; i-- {
        for _, v := range way2 {
            if way1[i] == v {
                return v
            }
        }
    }

    return nil
}
```



# 解法二：后序遍历

后续遍历很适合这种场合，先去左右子树找到 `p`, `q` 结点，看哪个结点的左右子树先找到左右结点则该结点就为答案。

```go
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
```

