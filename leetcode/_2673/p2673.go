package _2673

func minIncrements(n int, cost []int) (ans int) {
	idx := n - 1
	// 同步叶子结点中兄弟结点的cost
	for ; idx > n/2; idx -= 2 {
		maxVal := max(cost[idx], cost[idx-1])
		ans += 2*maxVal - cost[idx] - cost[idx-1]
		cost[idx], cost[idx-1] = maxVal, maxVal
	}
	// 同步非叶子结点的cost
	for ; idx > 0; idx -= 2 {
		myVal := cost[idx*2+1] + cost[idx]
		leftVal := cost[(idx-1)*2+1] + cost[idx-1]
		maxVal := max(myVal, leftVal)
		ans += 2*maxVal - myVal - leftVal
		cost[idx], cost[idx-1] = maxVal, maxVal
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
