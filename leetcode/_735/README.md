> [LeetCode 735 Asteroid Collision](https://leetcode.cn/problems/asteroid-collision)

# 解法：模拟，使用栈

```go
func asteroidCollision(asteroids []int) []int {
    var s []int
    for _, v := range asteroids {
        if v < 0 {
            // 把正向的“小”行星全部撞烂
            for len(s) > 0 && s[len(s) - 1] > 0 && s[len(s) - 1] < -v {
                s = s[:len(s) - 1]
            }
            // 如果栈里还有行星且为正向的，则这颗反向行星注定要爆炸，就无须执行下面的 append 了。
            if len(s) > 0 && s[len(s) - 1] > 0 {
                if s[len(s) - 1] == -v {
                    s = s[:len(s) - 1]
                }
                continue
            }
        }
        s = append(s, v)
    }
    return s
}
```

- 时间复杂度：$O(n)$
- 空间复杂度：$O(n)$
