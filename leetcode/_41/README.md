> [LeetCode 41 First Missing Positive](https://leetcode.cn/problems/first-missing-positive)

# 解法：原地哈希

将每个正数映射到数组对应的索引上，比如说 `1` 映射到索引 `0` 上，`2` 映射到索引 `1` 上。

也就是 `nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]`。

遍历映射完成后的数组，如果发现 `nums[i] != i + 1` ，则 `i + 1` 即为缺失的第一个正数。

完整代码如下：

```go
func firstMissingPositive(nums []int) int {
	n := len(nums)
    for i := 0; i < n; {
        if nums[i] != i + 1 && nums[i] > 0 && nums[i] - 1 < n && nums[nums[i] - 1] != nums[i] {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
			continue
        }
		i++
    }

	for idx, val := range nums {
		if val != idx + 1 {
			return idx + 1
		}
	}
	return n + 1
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(1)$
