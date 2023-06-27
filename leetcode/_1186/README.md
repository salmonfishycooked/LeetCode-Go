> [LeetCode 1186 Maximum Subarray Sum with One Deletion](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion)

# 记忆化搜索

首先找好状态，这里我找的是 `dfs(i, state)` ，代表数组 `arr` 中以第 `i` 号元素结尾的所有子数组中，最多删去 `state` 个数（当然可以不删）得到的最大和（但是至少保证选一个数）。

所以最后的答案就是 `max(dfs(0, 1), dfs(1, 1), ..., dfs(n - 1, 1))` 。

1. 如果 `state` 等于 `0` ，代表不能再删除了，所以 `dfs(i, 0) = max(dfs(i - 1, 0) + arr[i], arr[i])`。意思是要不我就在这里截断，要不我就再去 `i - 1` 那边看看。
2. 如果 `state` 等于 `1` ，代表最多可以删除 `1` 次，所以 `dfs(i, 1) = max(dfs(i - 1, 0), dfs(i - 1, 1) + arr[i], arr[i])`。意思是要不我在这里截断，要不我就删除这个数再去 `i - 1` 那边看看（`dfs(i - 1, 0)`），要不就不删，我再去 `i - 1` 那边看看要是能删一次能得到多大的和（`dfs(i - 1, 1 + arr[i])`）。

完整代码如下

```go
func maximumSum(arr []int) (res int) {
	const inf = 0x3f3f3f3f
	res = -inf
	dp := make([][2]int, len(arr))
	for i := range dp {
		dp[i] = [2]int{-inf, -inf}
	}
	var dfs func(i, state int) int
	dfs = func(i, state int) (res int) {
		if i < 0 {
			return -inf
		}
		p := &dp[i][state]
		if *p != -inf {
			return *p
		}
		defer func() { *p = res }()
		if state == 0 {
			return max(dfs(i-1, 0), 0) + arr[i]
		}
		return max(dfs(i-1, 0), max(dfs(i-1, 1)+arr[i], arr[i]))
	}
	for i := range arr {
		res = max(res, dfs(i, 1))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

- 时间复杂度：$O(n)$
- 空间复杂度：$O(n)$



## 翻译成递推

```go
func maximumSum(arr []int) (res int) {
    const inf = 0x3f3f3f3f
    res = -inf
    dp := make([][2]int, len(arr) + 1)
    dp[0] = [2]int{-inf, -inf}
    for i := range arr {
        dp[i + 1][0] = max(dp[i][0], 0) + arr[i]
        dp[i + 1][1] = max(dp[i][0], max(dp[i][1] + arr[i], arr[i]))
        res = max(res, dp[i + 1][1])
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

- 时间复杂度：$O(n)$
- 空间复杂度：$O(n)$



## 空间优化

```go
func maximumSum(arr []int) (res int) {
    const inf = 0x3f3f3f3f
    res = -inf
    f0, f1 := -inf, -inf
    for _, v := range arr {
        f1 = max(f0, max(f1 + v, v))
        f0 = max(f0, 0) + v
        res = max(res, f1)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

- 时间复杂度：$O(n)$
- 空间复杂度：$O(1)$