package _46

func permute(nums []int) (ans [][]int) {
	n := len(nums)
	set := make([]int, n)
	var dfs func(i, msk int)
	dfs = func(i, msk int) {
		if i == n {
			ans = append(ans, append([]int{}, set...))
			return
		}
		for j, v := range nums {
			if (msk>>j)&1 == 1 {
				set[i] = v
				dfs(i+1, msk^(1<<j))
			}
		}
	}
	dfs(0, (1<<n)-1)
	return
}
