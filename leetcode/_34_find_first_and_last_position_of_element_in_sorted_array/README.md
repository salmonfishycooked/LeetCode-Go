# [34. Find First and Last Position of Element in Sorted Array](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array)

Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

 

**Example 1:**

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

**Example 2:**

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

**Example 3:**

```
Input: nums = [], target = 0
Output: [-1,-1]
```

 

**Constraints:**

- `0 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`
- `nums` is a non-decreasing array.
- `-109 <= target <= 109`



## 解法一：二分查找

二分查找通用思路详见：LeetCode 162 README.md

完整代码如下

```go
func searchRange(nums []int, target int) []int {
	start := binarySearch(nums, target) + 1
	end := binarySearch(nums, target+1)
	if start > end {
		return []int{-1, -1}
	}
	return []int{start, end}
}

// binarySearch 会返回非递减数组 nums 中小于 target 的最大索引
func binarySearch(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 != r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
```

