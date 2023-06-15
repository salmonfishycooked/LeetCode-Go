package _209

import "math"

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
