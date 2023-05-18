package _2685

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
		// 如果该点未被访问过，则搜索该点
		if !vis[i] {
			v, e = 0, 0
			dfs(i)
			if e == v*(v-1) {
				ans++
			}
		}
	}
	return
}
