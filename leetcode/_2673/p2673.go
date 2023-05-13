package _2673

func minIncrements(n int, cost []int) (ans int) {
	// 同步叶子结点中兄弟结点的cost
	for i := n / 2; i < n; i += 2 {
		maxVal := max(cost[i], cost[i+1])
		ans += 2*maxVal - cost[i] - cost[i+1]
		cost[i], cost[i+1] = maxVal, maxVal
	}
	// 同步非叶子结点的cost
	idx := n/2 - 1
	for idx > 0 {
		// 上层结点索引
		upperLevel := idx/2 - 1
		// 将当前层结点都同步一下cost，使其路径相等
		for idx != upperLevel {
			myVal := cost[idx*2+1] + cost[idx]
			leftVal := cost[(idx-1)*2+1] + cost[idx-1]
			maxVal := max(myVal, leftVal)
			ans += 2*maxVal - myVal - leftVal
			cost[idx], cost[idx-1] = maxVal, maxVal
			idx -= 2
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
