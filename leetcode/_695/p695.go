package _695

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
