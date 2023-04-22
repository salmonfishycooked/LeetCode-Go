package _1_Two_Sum

// twoSum
// 一般思路，暴力枚举
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
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

// twoSum2
// 哈希映射解法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{i, idx}
		}
		m[v] = i
	}
	return nil
}
