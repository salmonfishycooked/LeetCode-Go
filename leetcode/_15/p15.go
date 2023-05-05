package _15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for k, v := range nums {
		if k > 0 && nums[k-1] == v {
			continue
		}
		right := n - 1
		for left := k + 1; left < right; left++ {
			if left > k+1 && nums[left-1] == nums[left] {
				continue
			}
			for left < right && nums[left]+nums[right]+v > 0 {
				right--
			}
			if left < right && nums[left]+nums[right]+v == 0 {
				ans = append(ans, []int{nums[left], nums[right], v})
			}
		}
	}
	return ans
}
