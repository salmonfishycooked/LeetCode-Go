package _1010

func numPairsDivisibleBy60(time []int) (cnt int) {
	// 记录歌曲对60取余后得到的数的个数
	m := make(map[int]int)
	for _, v := range time {
		// 如果在该数之前有余数与该数的余数之和等于60（也就是 60-v%60 )
		// 这里再对里层取余是防止v本来就是60的倍数
		if c, ok := m[(60-v%60)%60]; ok {
			cnt += c
		}
		m[v%60]++
	}
	return
}
