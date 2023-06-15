> [LeetCode 6455][https://leetcode.cn/problems/minimum-cost-to-make-all-characters-equal]

# 动态规划

记 `f[i]` 为使字符串 `0 ~ i - 1 ` 中所有字符相等的最小成本，

1. 若 `s[i] == s[i - 1]` ，则 `f[i] = f[i - 1]`。

2. 若 `s[i] != s[i - 1]` ，则 `f[i] = f[i - 1] + min(i, n - i)`。

​		`min(i, n - i)` 的意思是要使得第 `i` 个字符与前面的字符相等的话，要不就是翻转前 `i - 1` 个字符，要不		就是翻转包括 `i` 以后的所有字符，两者所需要的代价分别为 `i` 和 `n - i` ，根据 `f` 的定义需要求出所有可能		的最小值。（`n` 为字符串长度）

​		最后根据题意需要返回 `f[n - 1]` 。

完整代码如下：

```go
func minimumCost(s string) int64 {
	var ans int64
	for i, n := 1, len(s); i < n; i++ {
		if s[i] != s[i-1] {
			ans += int64(min(i, n-i))
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(1)$