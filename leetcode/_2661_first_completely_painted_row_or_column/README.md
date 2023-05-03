# [2661. First Completely Painted Row or Column](https://leetcode.cn/problems/first-completely-painted-row-or-column)

You are given a **0-indexed** integer array `arr`, and an `m x n` integer **matrix** `mat`. `arr` and `mat` both contain **all** the integers in the range `[1, m * n]`.

Go through each index `i` in `arr` starting from index `0` and paint the cell in `mat` containing the integer `arr[i]`.

Return *the smallest index* `i` *at which either a row or a column will be completely painted in* `mat`.

 

**Example 1:**

![image explanation for example 1](./assets/grid1.jpg)

```
Input: arr = [1,3,4,2], mat = [[1,4],[2,3]]
Output: 2
Explanation: The moves are shown in order, and both the first row and second column of the matrix become fully painted at arr[2].
```

**Example 2:**

![image explanation for example 2](./assets/grid2.jpg)

```
Input: arr = [2,8,7,4,1,3,5,6,9], mat = [[3,2,5],[1,4,6],[8,7,9]]
Output: 3
Explanation: The second column becomes fully painted at arr[3].
```

 

**Constraints:**

- `m == mat.length`
- `n = mat[i].length`
- `arr.length == m * n`
- `1 <= m, n <= 105`
- `1 <= m * n <= 105`
- `1 <= arr[i], mat[r][c] <= m * n`
- All the integers of `arr` are **unique**.
- All the integers of `mat` are **unique**.



# 解法一：哈希表

用 `mp` 来标记元素是否被上色，`pos` 记录元素在矩阵中出现的位置（索引）。

```go
func firstCompleteIndex(arr []int, mat [][]int) int {
    pos := make(map[int][2]int)
    // 记录mat中每个元素出现的位置
    for i, _ := range mat {
        for j, v := range mat[i] {
            pos[v] = [2]int{i, j}
        }
    }

    mp := make(map[int]bool)
    for i, v := range arr {
        // 给值为v的单元格涂色
        mp[v] = true
        // 找到mat中v的位置
        x, y := pos[v][0], pos[v][1]
        // 判断行和列是否都被涂色
        flag1 := true
        for j := len(mat) - 1; j >= 0; j-- {
            if !mp[mat[j][y]] {
                flag1 = false
                break
            }
        }
        flag2 := true
        for j := len(mat[0]) - 1; j >= 0; j-- {
            if !mp[mat[x][j]] {
                flag2 = false
                break
            }
        }
        if flag1 || flag2 {
            return i
        }
    }
    return -1
}
```

> 时间复杂度：$O((m+n)^2)$
>
> 空间复杂度：$O(m + n)$



# 解法二：哈希表（换种思路优化一下）

不用 `mp` 来记录哪个色块被涂色，用 `row` 和 `col` 两个数组来记录当前行或当前列已经被涂色的单元格的数量，比如 `row[0]` 等于 2 就是说明矩阵的第 0 行已经被涂了两个色块。

完整的代码如下

```go
func firstCompleteIndex(arr []int, mat [][]int) int {
    pos := make(map[int][2]int)
    // 记录mat中每个元素出现的位置
    for i, _ := range mat {
        for j, v := range mat[i] {
            pos[v] = [2]int{i, j}
        }
    }

    m, n := len(mat), len(mat[0])
    // 用 row, col 记录当前行或当前列已经被涂色的单元格的数量
    row := make([]int, m)
    col := make([]int, n)
    for i, v := range arr {
        x, y := pos[v][0], pos[v][1]
        row[x]++
        col[y]++
        // 判断数v所在矩阵的行或列是否被涂满
        if row[x] == n || col[y] == m {
            return i
        }
    }
    return -1
}
```

> 时间复杂度：$O(m n)$
>
> 空间复杂度：$O(m + n)$

时间开销主要在第一次遍历矩阵记录每个元素出现的位置。