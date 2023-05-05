原题：https://leetcode.cn/problems/3sum

# 解法：排序 + 双指针

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for k, v := range nums {
		// 去重
		if k > 0 && nums[k-1] == v {
			continue
		}
		// 让右指针指向数组最后一个数
		right := n - 1
		for left := k + 1; left < right; left++ {
			// 去重
			if left > k+1 && nums[left-1] == nums[left] {
				continue
			}
			for left < right && nums[left]+nums[right]+v > 0 {
				right--
			}
			// left < right 防止指针指向同一个值导致加入不存在的三元组
			if left < right && nums[left]+nums[right]+v == 0 {
				ans = append(ans, []int{nums[left], nums[right], v})
			}
		}
	}
	return ans
}
```

