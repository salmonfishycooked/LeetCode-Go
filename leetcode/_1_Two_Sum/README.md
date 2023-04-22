# 1. [Two Sum](https://leetcode.cn/problems/two-sum/)

## 题目

Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`*.

You may assume that each input would have **exactly one solution**, and you may not use the *same* element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- `2 <= nums.length <= 104`

- `-109 <= nums[i] <= 109`

- `-109 <= target <= 109`

- **Only one valid answer exists.**



## 解法一：暴力枚举

首先对于这题我们很容易想到定一移一，所以给出的代码如下：

```go
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := 0; j < i; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```

> 时间复杂度：$O(n^2)$
>
> 空间复杂度：$O(1)$

最坏情况下数组中任两个元素都会被匹配一次



## 解法二：哈希映射⭐

### 想一想解法一可以怎么优化

对于解法一，我们仔细想一想为什么要两次遍历，这是因为我们的 `i` 在指定一个数的时候，我们不清除在索引 `i  ` 前面的数组是否含有 `key` 值（就是解法一中的 `key`，这个 `key` 能使得这 `nums[i] + key = target`），也就是不知道（`nums[0] ~ nums[i-1]`）中是否有这个 `key` 值，所以才需要再遍历一次前 `i` 个元素。这里我们可以想到一个数据结构——哈希映射，刚好可以解决我们所提出的这个问题。



首先需要一个整数映射到整数的哈希映射，这里我们需要它从值映射到索引。

```java
m := map[int]int{}
```

然后进行数组的遍历，当遍历到第 `i` 个元素的时候，我们需要先检测一下在 `i` 之前是否存在一个数等于`target - nums[i]`，也就是在哈希映射中查询 `target - nums[i]`，如果真的存在这个数，那么就返回它们的索引。

```java
// 如果在哈希映射中找到这个数，就返回它们的索引
if idx, ok := m[target-v]; ok {
    return []int{i, idx}
}
```

如果不存在，我们就需要将当前这个数（也就是 `nums[i]`）存到哈希映射中，然后继续遍历。

```java
// 存入哈希映射
m[v] = i
```



完整代码如下

```go
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{i, idx}
		}
		m[v] = i
	}
	return nil
}
```

> 时间复杂度：$O(n)$
>
> 空间复杂度：$O(n)$

其中空间开销主要在**哈希映射**上
