> [LeetCode 139][]

# 动态规划

这里规定 `f[i]` 的含义是字符串 `s` 的前 `[0, i - 1]` 个字符所对应的字符串能否被字典拼接出来。

```go
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for _, v := range wordDict {
			if i-len(v) >= 0 && f[i-len(v)] && strings.Contains(s[i-len(v):i], v) {
				f[i] = true
			}
		}
	}
	return f[n]
}
```