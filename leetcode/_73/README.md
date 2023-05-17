> [LeetCode 73][https://leetcode.cn/problems/set-matrix-zeroes/]

# 解法

我们在搜索矩阵中原有的 `0` 的时候，用原矩阵的第 `0` 行，第 `0` 列分别记录需要归零的行和列。

当然，这样做之前我们需要先将第 `0 ` 行，第 `0` 列的信息先记录下来。（记录它们原先是否有0）

```go
func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    colZero, rowZero := false, false
    for i := 0; i < n; i++ {
        if matrix[0][i] == 0 {
            rowZero = true
        }
    }
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            colZero = true
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }

    for i := 1; i < n; i++ {
        if matrix[0][i] == 0 {
            for j := 0; j < m; j++ {
                matrix[j][i] = 0
            }
        }
    }
    for i := 1; i < m; i++ {
        if matrix[i][0] == 0 {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }

    if rowZero {
        for i := 0; i < n; i++ {
            matrix[0][i] = 0
        }
    }
    if colZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
```

> 时间复杂度：$O(m \times n)$
>
> 空间复杂度：$O(1)$