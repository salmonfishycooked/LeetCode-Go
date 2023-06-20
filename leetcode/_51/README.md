> [LeetCode 51 N 皇后](https://leetcode.cn/problems/n-queens)

# 解法：回溯

全排列问题，减枝优化。

`valid` 函数用来检测以点 `(i, j)` 为中心的斜对角线上是否有其他皇后。

```go
func solveNQueens(n int) (ans [][]string) {
    row := make([]int, n)

    valid := func(i, j int) bool {
        for k := 0; k < i; k++ {
            if i + j == k + row[k] || i - j == k - row[k] {
                return false
            }
        }
        return true
    }

    var dfs func(i, msk int)
    dfs = func(i, msk int) {
        if i == n {
            set := make([]string, n)
            for i, v := range row {
                set[i] = strings.Repeat(".", v) + "Q" + strings.Repeat(".", n - v - 1)
            }
            ans = append(ans, set)
            return
        }
        for j := 0; j < n; j++ {
            if (msk >> j) & 1 == 1 && valid(i, j) {
                row[i] = j
                dfs(i + 1, msk ^ (1 << j))
            }
        }
    }
    dfs(0, (1 << n) - 1)
    return
}
```

- 时间复杂度：$O(n^2 \times n!)$。搜索树中至多有 $O(n!)$ 个叶子，每个叶子生成答案每次需要 $O(n^2)$ 的时间，所以时间复杂度为 $O(n^2 \times n!)$。实际上搜索树中远没有这么多叶子，`n = 9` 时只有 `352` 种放置方案，远远小于 `9! = 362880` 。更加准确的方案数可以参考 [OEIS A000170](https://leetcode.cn/link/?target=https%3A%2F%2Foeis.org%2FA000170)，为$O(\frac{n!}{2.54^n})$。
- 空间复杂度：$O(n)$，返回值所占的空间不计在内。



可以将 `valid` 函数用两个表示对角线情况的数组代替，降低每次判断 `valid` 所耗费的时间。

```go
func solveNQueens(n int) (ans [][]string) {
    row := make([]int, n)
    diag1 := make([]bool, 2*n - 1)
    diag2 := make([]bool, 2*n - 1)

    var dfs func(i, msk int)
    dfs = func(i, msk int) {
        if i == n {
            set := make([]string, n)
            for i, v := range row {
                set[i] = strings.Repeat(".", v) + "Q" + strings.Repeat(".", n - v - 1)
            }
            ans = append(ans, set)
            return
        }
        for j := 0; j < n; j++ {
            if (msk >> j) & 1 == 1 && !diag1[i + j] && !diag2[i - j + n - 1] {
                row[i] = j
                diag1[i + j], diag2[i - j + n - 1] = true, true
                dfs(i + 1, msk ^ (1 << j))
                diag1[i + j], diag2[i - j + n - 1] = false, false
            }
        }
    }
    dfs(0, (1 << n) - 1)
    return
}
```

时间复杂度和空间复杂度等同于上面。