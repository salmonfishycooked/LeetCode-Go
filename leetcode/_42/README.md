> [LeetCode 42 接雨水](https://leetcode.cn/problems/trapping-rain-water)

# 解法：双指针

```go
func trap(height []int) int {
    n := len(height)
    left, right := 0, n - 1
    leftMax, rightMax := 0, 0
    res := 0
    for left <= right {
        leftMax = max(leftMax, height[left])
        rightMax = max(rightMax, height[right])
        if leftMax < rightMax {
            res += leftMax - height[left]
            left++
        } else {
            res += rightMax - height[right]
            right--
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
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