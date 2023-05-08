>  [LeetCode 695](https://leetcode.cn/problems/max-area-of-island)

# 解法：深度优先搜索

## 递归写法

```go
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				ans = max(ans, dfs(grid, i, j))
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i int, j int) int {
	// 搜索越界或者访问的该位置不是土地
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0 // 标记该地已搜索
	ans := 1
	di := [4]int{0, 0, 1, -1}
	dj := [4]int{1, -1, 0, 0}
	for k := 0; k < 4; k++ {
		ans += dfs(grid, i+di[k], j+dj[k])
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```

> 时间复杂度：$O(mn)$
>
> 空间复杂度：$O(mn)$



## 非递归写法

```go
func maxAreaOfIsland(grid [][]int) int {
    ans := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] != 0 {
                curAns := 0
                // 用栈保存当前需要搜索的点
                var stack [][2]int
                stack = append(stack, [2]int{i, j})
                for len(stack) != 0 {
                    nowi, nowj := stack[len(stack)-1][0], stack[len(stack)-1][1]
                    stack = stack[:len(stack)-1]
                    // 如果该点越界或者该点不是土地就不用搜了
                    if !isSafe(grid, nowi, nowj) {
                        continue
                    }

                    // 该点是土地的话
                    grid[nowi][nowj] = 0 // 标记已搜过
                    curAns++ // 连通的土地数+1
                    di := [4]int{0, 0, 1, -1}
                    dj := [4]int{1, -1, 0, 0}
                    for k := 0; k < 4; k++ {
                        stack = append(stack, [2]int{nowi+di[k], nowj+dj[k]})
                    }
                }
                ans = max(ans, curAns)
            }
        }
    }
    return ans
}

func isSafe(grid [][]int, i, j int) bool {
    if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
        return false
    }
    return true
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
```

> 时间复杂度：$O(mn)$
>
> 空间复杂度：$O(mn)$



# 广度优先搜索

```go
func maxAreaOfIsland(grid [][]int) int {
    ans := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] != 0 {
                curAns := 0
                var queue [][2]int // 用队列保存当前需要搜索的点
                queue = append(queue, [2]int{i, j})
                for len(queue) != 0 {
                    nowi, nowj := queue[0][0], queue[0][1]
                    queue = queue[1:]
                    // 判断取出来的点是否越界或者该点不是土地
                    if !isSafe(grid, nowi, nowj) {
                        continue
                    }

                    curAns++
                    grid[nowi][nowj] = 0 // 标记该点已搜过
                    di := [4]int{0, 0, 1, -1}
                    dj := [4]int{1, -1, 0, 0}
                    for k := 0; k < 4; k++ {
                        queue = append(queue, [2]int{nowi+di[k], nowj+dj[k]})
                    }
                }
                ans = max(ans, curAns)
            }
        }
    }
    return ans
}

func isSafe(grid [][]int, i, j int) bool {
    if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
        return false
    }
    return true
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
```

> 时间复杂度：$O(mn)$
>
> 空间复杂度：$O(mn)$