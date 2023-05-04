package _128

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
