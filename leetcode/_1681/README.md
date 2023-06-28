> [LeetCode 1681 Minumum Incompatibility](https://leetcode.cn/problems/minimum-incompatibility)

# 状压DP

```go
func minimumIncompatibility(nums []int, k int) (res int) {
	n := len(nums)
	gsize := n / k
	sort.Ints(nums)

	// 如果数组中的某个数的数量多于组数，则划分的组一定不合法。
	for i, cnt := 0, 1; i < n-1; i++ {
		for i < n-1 && nums[i] == nums[i+1] {
			cnt++
			i++
		}
		if cnt > k {
			return -1
		}
		cnt = 1
	}

	f := make([][]int, 1<<n-1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	// rset 当前可选数的集合
	// pre 上一个选的数的下标
	var dfs func(rset, pre int) int
	dfs = func(rset, pre int) (res int) {
		if rset == 0 {
			return 0
		}
		p := &f[rset][pre]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		// 如果凑好一组了，重新凑组
		if bits.OnesCount(uint(rset))%gsize == 0 {
            // 取集合的最右边的1
			one := rset & (^rset + 1)
			return dfs(rset^one, bits.Len(uint(one))-1)
		}
		// 没凑好组，继续选数凑组
		res = math.MaxInt32
		for i := pre + 1; i < n; i++ {
			if rset>>i&1 == 1 && nums[i] != nums[pre] {
				res = min(res, nums[i]-nums[pre]+dfs(rset^(1<<i), i))
			}
		}
		return
	}
	return dfs(1<<n-2, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

- 时间复杂度：$O(n^2 \times 2^n)$。状态个数 $O(n \times 2^n)$，每个状态的处理时间 $O(n)$。

- 空间复杂度：$O(n \times 2^n)$。主要是记忆化数组所占空间。