package _1330

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
