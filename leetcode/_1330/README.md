> [LeetCode 1330][https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value] 🌟

# 解法一：暴力解法

我们可以通过观察发现，翻转其中任意一个子数组（假设该子数组左端点为 `i` ，右端点为 `j` ），得到的**数组值**的变化仅仅是子数组的端点处有变化，子数组内部差值的绝对值还是一样的。也就是翻转后得到的**数组值**为

$newDiff = sumDiff - abs(nums[i] - nums[i - 1]) - abs(nums[j] - nums[j + 1]) + abs(nums[j] - nums[i - 1]) + abs(nums[i] - nums[j + 1])$

上述情况只适用于子数组的端点不与原数组的端点重合的情况，重合的情况我们另算即可。

这样，我们只需枚举所有子数组就能得到答案。

```go
func maxValueAfterReverse(nums []int) (ans int) {
    n := len(nums)
    sumDiff := 0
    for i := 1; i < n; i++ {
        sumDiff += abs(nums[i] - nums[i - 1])
    }
    ans = sumDiff

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            var curSum int
            if i > 0 {
                if j < n - 1 {
                    curSum = sumDiff - abs(nums[i] - nums[i - 1]) - abs(nums[j] - nums[j + 1]) + abs(nums[j] - nums[i - 1]) + abs(nums[i] - nums[j + 1])
                } else {
                    curSum = sumDiff - abs(nums[i] - nums[i - 1]) + abs(nums[j] - nums[i - 1])
                }
            } else if i == 0 {
                if j < n - 1 {
                    curSum = sumDiff - abs(nums[j] - nums[j + 1]) + abs(nums[i] - nums[j + 1])
                } else {
                    curSum = sumDiff
                }
            }
            ans = max(ans, curSum)
        }
    }
    return
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

> 时间复杂度：$O(n^2)$
>
> 空间复杂度：$O(1)$

毫无疑问，根据题目的数据量来看，会超时。



# 解法二：贪心

感谢题解：https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/solutions/1680613/chun-cui-by-bai-mu-ying-li-dra-eniac-rd49/ 提供的思路。

这里给出 Go 的实现。

```go
func maxValueAfterReverse(nums []int) int {
	n := len(nums)
	sumDiff := 0
	for i := 1; i < n; i++ {
		sumDiff += abs(nums[i] - nums[i-1])
	}

	delta := 0
	delta = max(delta, calLeft(nums))
	delta = max(delta, calRight(nums))
	delta = max(delta, calMiddle(nums))
	return sumDiff + delta
}

// calLeft 翻转的子数组左端点为0的情况
func calLeft(nums []int) (v int) {
	for right, n := 1, len(nums); right < n-1; right++ {
		v = max(v, abs(nums[0]-nums[right+1])-abs(nums[right]-nums[right+1]))
	}
	return v
}

// calRight 翻转的子数组右端点为 n - 1 的情况
func calRight(nums []int) (v int) {
	for left, n := 1, len(nums); left < n-1; left++ {
		v = max(v, abs(nums[n-1]-nums[left-1])-abs(nums[left]-nums[left-1]))
	}
	return v
}

func calMiddle(nums []int) int {
	pairLessMax := min(nums[0], nums[1])
	pairMoreMin := max(nums[0], nums[1])
	for i, n := 1, len(nums); i < n-1; i++ {
		pairLessMax = max(pairLessMax, min(nums[i], nums[i+1]))
		pairMoreMin = min(pairMoreMin, max(nums[i], nums[i+1]))
	}
	return 2 * (pairLessMax - pairMoreMin)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
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
