package _139

import "strings"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for _, v := range wordDict {
			if i-len(v) >= 0 && f[i-len(v)] && strings.Contains(s[i-len(v):i], v) {
				f[i] = true
			}
		}
	}
	return f[n]
}
