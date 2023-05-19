> [LeetCode 1079][https://leetcode.cn/problems/letter-tile-possibilities]

# 解法：回溯

```go
func numTilePossibilities(tiles string) (ans int) {
	n := len(tiles)
	vis := make([]bool, n)
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			return
		}
		// 在当前层不能选择同样的字母，所以用哈希表来记录已经选过的字母
		// 用数组同理，因为该题字符集仅为所有大写英文字母
		m := make(map[byte]bool)
		for i := 0; i < n; i++ {
			if vis[i] || m[tiles[i]] {
				continue
			}
			vis[i], m[tiles[i]] = true, true
			ans++
			dfs(pos + 1)
			vis[i] = false
		}
	}
	dfs(0)
	return
}
```

> 时间复杂度：$O(n \times n!)$
>
> 空间复杂度：$O(n)$

时间最坏情况下：`tiles` 字符串里面没有重复字母，$dfs$ 里面会有 $n!$ 种状态，每层都会被线性扫描一遍。

空间最坏情况下：`tiles` 字符串里面没有重复字母，调用栈的深度为 `n` ，每层仅用有限的空间来存储数据。