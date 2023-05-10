> [LeetCode 46][https://leetcode.cn/problems/permutations]

# 解法：回溯

```go
func permute(nums []int) (ans [][]int) {
	n := len(nums)
	var backtrack func(int)
	tmp := make([]int, 0, n)
	vis := make([]bool, n)
	backtrack = func(level int) {
		if level == n {
			t := make([]int, n)
			copy(t, tmp)
			ans = append(ans, t)
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			tmp = append(tmp, nums[i])
			vis[i] = true
			backtrack(level + 1)
			vis[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(0)
	return
}

```

