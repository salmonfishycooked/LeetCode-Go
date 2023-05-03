package _1003_check_if_word_is_valid_after_substitutions

func isValid(s string) bool {
	var stack []byte
	for _, s := range s {
		if l := len(stack); s == 'c' {
			if l < 2 {
				return false
			}
			if stack[l-1] == 'b' && stack[l-2] == 'a' {
				stack = stack[:l-2]
				continue
			}
		}
		stack = append(stack, byte(s))
	}
	return len(stack) == 0
}
