# [1003.Check If Word Is Valid After Substitutions](https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions)

Given a string `s`, determine if it is **valid**.

A string `s` is **valid** if, starting with an empty string `t = ""`, you can **transform** `t` **into** `s` after performing the following operation **any number of times**:

- Insert string `"abc"` into any position in `t`. More formally, `t` becomes `tleft + "abc" + tright`, where `t == tleft + tright`. Note that `tleft` and `tright` may be **empty**.

Return `true` *if* `s` *is a **valid** string, otherwise, return* `false`.

 

**Example 1:**

```
Input: s = "aabcbc"
Output: true
Explanation:
"" -> "abc" -> "aabcbc"
Thus, "aabcbc" is valid.
```

**Example 2:**

```
Input: s = "abcabcababcc"
Output: true
Explanation:
"" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
Thus, "abcabcababcc" is valid.
```

**Example 3:**

```
Input: s = "abccba"
Output: false
Explanation: It is impossible to get "abccba" using the operation.
```

 

**Constraints:**

- `1 <= s.length <= 2 * 104`
- `s` consists of letters `'a'`, `'b'`, and `'c'`



# 解法：栈

遇到字符 `c` 的时候，依次从栈顶拿出两个元素，看是否能组成 `abc` ，不行的话直接返回 `false` ，最后还需要检查一下栈是否为空。

完整代码如下

```go
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
```

