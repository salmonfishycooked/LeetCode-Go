# [1010. Pairs of Songs With Total Durations Divisible by 60](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60)

You are given a list of songs where the `ith` song has a duration of `time[i]` seconds.

Return *the number of pairs of songs for which their total duration in seconds is divisible by* `60`. Formally, we want the number of indices `i`, `j` such that `i < j` with `(time[i] + time[j]) % 60 == 0`.

 

**Example 1:**

```
Input: time = [30,20,150,100,40]
Output: 3
Explanation: Three pairs have a total duration divisible by 60:
(time[0] = 30, time[2] = 150): total duration 180
(time[1] = 20, time[3] = 100): total duration 120
(time[1] = 20, time[4] = 40): total duration 60
```

**Example 2:**

```
Input: time = [60,60,60]
Output: 3
Explanation: All three pairs have a total duration of 120, which is divisible by 60.
```

 

**Constraints:**

- `1 <= time.length <= 6 * 104`
- `1 <= time[i] <= 500`



# 解法：哈希表

用一个长度为 60 的数组去存储也可以，这里哈希表所占的空间同样也是 $O(1)$ 级别。

```go
func numPairsDivisibleBy60(time []int) (cnt int) {
	// 记录歌曲对60取余后得到的数的个数
	m := make(map[int]int)
	for _, v := range time {
		// 如果在该数之前有余数与该数的余数之和等于60（也就是 60-v%60 )
		// 这里再对里层取余是防止v本来就是60的倍数
		if c, ok := m[(60-v%60)%60]; ok {
			cnt += c
		}
		m[v%60]++
	}
	return
}
```

