package _1079

func numTilePossibilities(tiles string) (ans int) {
	n := len(tiles)
	vis := make([]bool, n)
	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == n {
			return
		}
		// 在当前层不能选择同样的字母，所以用哈希表来记录已经选过的字母
		// 用数组同理，因为该题字符集仅为所有大写英文字母
		m := make(map[byte]bool)
		for i := 0; i < n; i++ {
			if vis[i] || m[tiles[i]] {
				continue
			}
			vis[i], m[tiles[i]] = true, true
			ans++
			backtrack(pos + 1)
			vis[i] = false
		}
	}
	backtrack(0)
	return
}
