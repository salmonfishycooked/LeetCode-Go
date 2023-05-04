# 128. Longest Consecutive Sequence

Given an unsorted array of integers `nums`, return *the length of the longest consecutive elements sequence.*

You must write an algorithm that runs in `O(n)` time.

 

**Example 1:**

```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

**Example 2:**

```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

 

**Constraints:**

- `0 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`



# 解法：哈希表

```go
func longestConsecutive(nums []int) int {
	// 记录数组中出现过的数
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	ans := 0
	for key, _ := range m {
		// 若该数前面不存在数，则把该数当作连续序列的第一个数
		if !m[key-1] {
			long := 1
			for m[key+1] {
				key++
				long++
			}
			if long > ans {
				ans = long
			}
		}
	}
	return ans
}
```

