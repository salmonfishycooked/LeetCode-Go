package _2106

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	// sum[i]代表第i + 1个水果之前（不包括第i个）总共的水果数量
	// idx[i]代表第i + 1个水果所在的 position
	sum := make([]int, 0, n+1)
	idx := make([]int, 0, n)
	sum = append(sum, 0) // 第1个水果之前没有水果
	for i := 0; i < n; i++ {
		sum = append(sum, sum[i]+fruits[i][1])
		idx = append(idx, fruits[i][0])
	}

	ans := 0
	for step := 0; step <= k/2; step++ {
		// 先往左走
		start := startPos - step
		end := startPos + (k - 2*step)
		l := low(idx, 0, len(idx), start)
		r := upp(idx, 0, len(idx), end)
		if tmp := sum[r] - sum[l]; tmp > ans {
			ans = tmp
		}
		// 先往右走
		start = startPos - (k - 2*step)
		end = startPos + step
		l = low(idx, 0, len(idx), start)
		r = upp(idx, 0, len(idx), end)
		if tmp := sum[r] - sum[l]; tmp > ans {
			ans = tmp
		}
	}
	return ans
}

func low(arr []int, l int, r int, t int) int {
	l1, r1 := l-1, r
	for l1+1 != r1 {
		mid := l1 + (r1-l1)>>1
		if arr[mid] < t {
			l1 = mid
		} else {
			r1 = mid
		}
	}
	return r1
}

func upp(arr []int, l int, r int, t int) int {
	l1, r1 := l-1, r
	for l1+1 != r1 {
		mid := l1 + (r1-l1)>>1
		if arr[mid] <= t {
			l1 = mid
		} else {
			r1 = mid
		}
	}
	return r1
}
