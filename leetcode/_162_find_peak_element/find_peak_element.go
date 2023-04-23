package _162_find_peak_element

// findPeakElement
// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func findPeakElement(nums []int) int {
	l, r := -1, len(nums)-1
	for l+1 != r {
		mid := (r-l)/2 + l
		if nums[mid] < nums[mid+1] {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}
