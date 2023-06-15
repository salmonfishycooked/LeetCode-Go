> [LeetCode 209 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum)

# 解法：滑动窗口

使用同向双指针实现。

完整代码如下：

```go
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	res := math.MaxInt32
	for right, v := range nums {
		sum += v
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(1)$