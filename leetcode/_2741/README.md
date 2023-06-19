> [LeetCode 2741 特别的排列](https://leetcode.cn/problems/special-permutations)

# 解法：状压 DP

```go
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
```

> 时间复杂度：$O(n^2 \times 2^n)$
>
> 空间复杂度：$O(n \times 2^n)$

$时间复杂度$ = $O(状态个数) \times O(单个状态的计算时间)$

​					= $O(n \times 2^n) \times O(n)$

​					= $O(n^2 \times 2^n)$

空间复杂度即为数组 `memos` 所开辟的内存大小，递归调用栈所开辟的空间数量级远小于前者，忽略不计。
