package _1186

func maximumSum(arr []int) (res int) {
	const inf = 0x3f3f3f3f
	res = -inf
	f0, f1 := -inf, -inf
	for _, v := range arr {
		f1 = max(f0, max(f1+v, v))
		f0 = max(f0, 0) + v
		res = max(res, f1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
