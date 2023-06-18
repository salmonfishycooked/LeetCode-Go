package _41

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] != i+1 && nums[i] > 0 && nums[i]-1 < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			continue
		}
		i++
	}

	for idx, val := range nums {
		if val != idx+1 {
			return idx + 1
		}
	}
	return n + 1
}
