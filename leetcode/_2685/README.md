> [LeetCode 2685][https://leetcode.cn/problems/count-the-number-of-complete-components/]

在完全图中，任意两点都有边，相当于从 $v$ 个点中选 $2$ 个点的组合数，即

$C_{v}^{2} = \frac{v(v-1)}{2}$

为了统计方便，我们统计边的时候每个点的边都会统计，这就会让每条边被统计两次，所以最后的边数如果为

$e = v(v - 1)$ 即为完全图。

# 解法一：DFS

```go
func countCompleteComponents(n int, edges [][]int) (ans int) {
    g := make([][]int, n)
    // 建图
    for _, v := range edges {
        x, y := v[0], v[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    vis := make([]bool, n)
    
    var v, e int
    var dfs func(n int)
    dfs = func(n int) {
        vis[n] = true
        v++
        e += len(g[n])
        for _, v := range g[n] {
            if !vis[v] {
                dfs(v)
            }
        }
    }

    for i := 0; i < n; i++ {
        if !vis[i] {
            v, e = 0, 0
            dfs(i)
            if e == v*(v - 1) {
                ans++
            }
        }
    }
    return
}
```



# 解法二：BFS

```go
func countCompleteComponents(n int, edges [][]int) (ans int) {
    g := make([][]int, n)
    // 建图
    for _, v := range edges {
        x, y := v[0], v[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    vis := make([]bool, n)
    var v, e int
    for i := 0; i < n; i++ {
        var queue []int
        if !vis[i] {
            v, e = 0, 0
            queue = append(queue, i)
            vis[i] = true
            for len(queue) > 0 {
                node := queue[0]
                queue = queue[1:]
                v++
                e += len(g[node])
                for _, x := range g[node] {
                    if !vis[x] {
                        vis[x] = true
                        queue = append(queue, x)
                    }
                }
            }
            if e == v*(v - 1) {
                ans++
            }
        }
    }
    return
}
```

> 时间复杂度：$O(m + n)$
>
> 空间复杂度：$O(m + n)$

其中 $n$ 为 `edges` 的长度。

