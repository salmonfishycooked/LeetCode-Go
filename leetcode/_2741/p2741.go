package _2741

func specialPerm(nums []int) (ans int) {
	n := len(nums)
	m := 1 << n
	const mod int = 1e9 + 7
	memos := make([][]int, m)
	for i := range memos {
		memos[i] = make([]int, n)
		for j := range memos[i] {
			memos[i][j] = -1
		}
	}

	var dfs func(set, idx int) int
	dfs = func(set, idx int) (res int) {
		// 可选集为空集，找到一个特殊排列！
		if set == 0 {
			return 1
		}
		if memos[set][idx] != -1 {
			return memos[set][idx]
		}
		for i, v := range nums {
			if set>>i&1 == 1 && (v%nums[idx] == 0 || nums[idx]%v == 0) {
				res = (res + dfs(set^(1<<i), i)) % mod
			}
		}
		memos[set][idx] = res
		return
	}

	// m - 1 为全集
	for i := 0; i < n; i++ {
		ans = (ans + dfs((m-1)^(1<<i), i)) % mod
	}
	return
}
