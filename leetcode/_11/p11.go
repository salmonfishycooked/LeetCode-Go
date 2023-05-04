package _11

func maxArea(height []int) int {
	l := len(height)
	left, right := 0, l-1
	ans := 0
	for left != right {
		if tmp := (right - left) * min(height[left], height[right]); tmp > ans {
			ans = tmp
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
