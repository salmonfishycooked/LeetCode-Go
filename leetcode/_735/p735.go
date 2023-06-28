package _735

func asteroidCollision(asteroids []int) []int {
	var s []int
	for _, v := range asteroids {
		if v < 0 {
			for len(s) > 0 && s[len(s)-1] > 0 && s[len(s)-1] < -v {
				s = s[:len(s)-1]
			}
			if len(s) > 0 && s[len(s)-1] > 0 {
				if s[len(s)-1] == -v {
					s = s[:len(s)-1]
				}
				continue
			}
		}
		s = append(s, v)
	}
	return s
}
