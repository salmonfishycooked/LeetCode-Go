package _6455

func minimumCost(s string) int64 {
	var ans int64
	for i, n := 1, len(s); i < n; i++ {
		if s[i] != s[i-1] {
			ans += int64(min(i, n-i))
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
