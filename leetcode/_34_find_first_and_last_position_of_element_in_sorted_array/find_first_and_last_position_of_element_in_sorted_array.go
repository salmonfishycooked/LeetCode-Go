package _34_find_first_and_last_position_of_element_in_sorted_array

// searchRange
// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
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
