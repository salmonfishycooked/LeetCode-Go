# [11. Container With Most Water](https://leetcode.cn/problems/container-with-most-water)

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return *the maximum amount of water a container can store*.

**Notice** that you may not slant the container.

 

**Example 1:**

![img](./assets/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
```

**Example 2:**

```
Input: height = [1,1]
Output: 1
```

 

**Constraints:**

- `n == height.length`
- `2 <= n <= 105`
- `0 <= height[i] <= 104`



# 解法：双指针

官方的题解讲得挺详细的，这里给出 Go 的实现。

```go
func maxArea(height []int) int {
    l := len(height)
    left, right := 0, l - 1
    ans := 0
    for left != right {
        if tmp := (right - left) * min(height[left], height[right]); tmp > ans {
            ans = tmp
        }
        if height[left] > height[right] {
            right--
        } else {
            left++
        }
    }
    return ans
}

func min(a int, b int) int {
    if a > b {
        return b
    }
    return a
}
```

