>  [LeetCode 1263][https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location]

# 解法：双端队列BFS

```go
func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var si, sj, bi, bj int

	// 找出玩家的位置(si, sj)和箱子的位置(bi, bj)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'B' {
				bi, bj = i, j
			}
			if grid[i][j] == 'S' {
				si, sj = i, j
			}
		}
	}

	// 用于将二维状态进行压缩的函数
	f := func(i, j int) int {
		return i*n + j
	}

	// 用于检查位置(i, j)是否越界或者是否为墙
	isSafe := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != '#'
	}

	q := [][]int{[]int{f(si, sj), f(bi, bj), 0}}
	vis := make([][]bool, m*n)
	for i := range vis {
		vis[i] = make([]bool, m*n)
	}
	vis[f(si, sj)][f(bi, bj)] = true
	dir := [5]int{-1, 0, 1, 0, -1}

	for len(q) > 0 {
		// 出队
		state := q[0]
		q = q[1:]
		// 取出状态对应的信息
		si, sj, bi, bj = state[0]/n, state[0]%n, state[1]/n, state[1]%n
		cnt := state[2]
		// 判断该状态下箱子是否在终点
		if grid[bi][bj] == 'T' {
			return cnt
		}
		// 若不在终点，玩家可以朝四个方向移动
		for k := 0; k < 4; k++ {
			sx, sy := si+dir[k], sj+dir[k+1]
			// 如果该点非安全点
			if !isSafe(sx, sy) {
				continue
			}
			// 如果是安全点，进行下一步
			// 玩家进行一次移动后，如果该位置是箱子的位置，箱子也会对应移动
			if sx == bi && sy == bj {
				// 箱子移动的方向与玩家移动的方向相同
				bx, by := bi+dir[k], bj+dir[k+1]
				if !isSafe(bx, by) || vis[f(sx, sy)][f(bx, by)] {
					continue
				}
				vis[f(sx, sy)][f(bx, by)] = true
				q = append(q, []int{f(sx, sy), f(bx, by), cnt + 1})
			} else if !vis[f(sx, sy)][f(bi, bj)] {
				// 若移动后的玩家位置不处于箱子的位置
				vis[f(sx, sy)][f(bi, bj)] = true
				q = append([][]int{[]int{f(sx, sy), f(bi, bj), cnt}}, q...)
			}
		}
	}
	return -1
}
```

> 时间复杂度：$O(m^2 \times n^2)$
>
> 空间复杂度：$O(m^2 \times n^2)$