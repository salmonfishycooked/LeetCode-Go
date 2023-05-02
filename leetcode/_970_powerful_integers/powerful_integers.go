package _970_powerful_integers

func powerfulIntegers(x int, y int, bound int) []int {
	m := make(map[int]bool)
	v := 1
	for i := 0; i < 20; i++ {
		v1 := 1
		for j := 0; j < 20; j++ {
			v2 := v + v1
			if v2 <= bound {
				m[v2] = true
			} else {
				break
			}
			v1 *= y
		}
		if v > bound {
			break
		}
		v *= x
	}

	var ans []int
	for v := range m {
		ans = append(ans, v)
	}
	return ans
}
