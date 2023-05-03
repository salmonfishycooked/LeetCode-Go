package _1376_time_needed_to_inform_all_employees

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make(map[int][]int)
	// id为v的员工掌管着一个员工id数组
	for i, v := range manager {
		g[v] = append(g[v], i)
	}

	var dfs func(int) int
	// dfs 含义是以 id 为根结点的树，通知完其所有子节点所需要的时间
	dfs = func(id int) (childrenTime int) {
		for _, childID := range g[id] {
			time := dfs(childID)
			if time > childrenTime {
				childrenTime = time
			}
		}
		return informTime[id] + childrenTime
	}

	return dfs(headID)
}
