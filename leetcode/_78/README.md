> [LeetCode 78][https://leetcode.cn/problems/subsets/]

# 解法：回溯

```go
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	var tmp []int

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == n {
			trip := make([]int, len(tmp))
			copy(trip, tmp)
			ans = append(ans, trip)
			return
		}

		// 不选nums[pos]
		backtrack(pos + 1)

		// 选nums[pos]
		tmp = append(tmp, nums[pos])
		backtrack(pos + 1)

		// 撤销选择
		tmp = tmp[:len(tmp)-1]
	}
	backtrack(0)
	return
}
```

