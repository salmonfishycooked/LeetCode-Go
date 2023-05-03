package _49_group_anagrams

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, v := range strs {
		cnt := [26]int{}
		for _, c := range v {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], v)
	}

	ans := make([][]string, 0, len(strs))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
