package _1090

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})

	res, cnt := 0, numWanted
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		if v, ok := m[labels[idx[i]]]; ok && v == useLimit {
			continue
		}
		res += values[idx[i]]
		m[labels[idx[i]]]++
		cnt--
		if cnt == 0 {
			break
		}
	}
	return res
}
