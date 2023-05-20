> [LeetCode 72][https://leetcode.cn/problems/edit-distance/]

# 暴力枚举

我们定义函数 $f(i, j)$ 的作用是返回将字符串 `word1` 的前 `i` 个字符与 `word2` 的前 `j` 个字符相等的最少操作数。

假设我们已经到了最后一步，就是让 `word1[i]` 与 `word2[j]` 相等。

如果它们本来就是相等的，那我们编辑距离就是与 `f(i - 1, j - 1)` 相等的。

如果不同则有三种方法，最后的完全态由 `f(i, j - 1)` 经过插入得来，或者由 `f(i - 1, j)` 通过删除得来，或者由 `f(i - 1, j - 1)` 经过替换得来。

递归退出条件就是 `i == -1` 或者 `j == -1` 此时，对应的情况是，一条字符串为空串，所以应该删除另一条字符串的 [0, i - 1] 的所有字符，需要操作 `i + 1` （或`j + 1`）次。

```go
const INF = 0x3f3f3f3f

func minDistance(word1 string, word2 string) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if j == -1 {
			return i + 1
		}
		if i == -1 {
			return j + 1
		}

		if word1[i] == word2[j] {
			return f(i-1, j-1)
		}
		return min(min(f(i, j-1)+1, f(i-1, j)+1), f(i-1, j-1)+1)
	}
	return f(len(word1)-1, len(word2)-1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```





# 记忆化搜索

暴力枚举的时间复杂度太高了，是 3 的指数级别，这是因为我们计算了大量的重复子问题，所以我们缓存计算出的子问题的答案，如果命中缓存的话直接返回即可。这会将我们的时间复杂度降为 $O(m \times n)$，同时空间复杂度也会上升到 $O(m \times n)$。

```go
const INF = 0x3f3f3f3f

func minDistance(word1 string, word2 string) int {
    n1, n2 := len(word1), len(word2)
    mem := make([][]int, n1)
    for i := 0; i < n1; i++ {
        mem[i] = make([]int, n2)
        for j := 0; j < n2; j++ {
            mem[i][j] = INF
        }
    }

    var f func(i, j int) int
    f = func(i, j int) int {
        if j == -1 {
            return i + 1
        }
        if i == -1 {
            return j + 1
        }
        if mem[i][j] != INF {
            return mem[i][j]
        }

        if word1[i] == word2[j] {
            mem[i][j] = f(i - 1, j - 1)
        } else {
            mem[i][j] = min(min(f(i, j - 1) + 1, f(i - 1, j) + 1), f(i - 1, j - 1) + 1)
        }
        return mem[i][j]
    }
    return f(len(word1) - 1, len(word2) - 1)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```



# 动态规划

有了前两个方法做铺垫，写出动态规划也就信手拈来了。

```go
const INF = 0x3f3f3f3f

func minDistance(word1 string, word2 string) int {
    n1, n2 := len(word1), len(word2)
    f := make([][]int, n1 + 1)
    for i := 0; i <= n1; i++ {
        f[i] = make([]int, n2 + 1)
    }
    for i := 0; i <= n1; i++ {
        f[i][0] = i
    }
    for i := 0; i <= n2; i++ {
        f[0][i] = i
    }

    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            if word1[i - 1] == word2[j - 1] {
                f[i][j] = f[i - 1][j - 1]
                continue
            }
            f[i][j] = min(min(f[i][j - 1], f[i - 1][j]), f[i - 1][j - 1]) + 1
        }
    }
    return f[n1][n2]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```



**动态规划的本质就是，暴力枚举所有可能（dfs整颗枚举树）+ 记忆化。**

1. 我们首先先想出暴力枚举的递归方法（使用分治思想，假设我这个小问题已经解决了，现在我要用这个小问题来解决大问题，在这之前要先明确递归函数的返回值的意义），然后我们处理细节，比如说递归退出条件（这个其实就是动态规划的 base case）。
2. 再将步骤1写出来的递归进行记忆化，变成记忆化搜索
3. 将递归自顶向下解决问题的方法改写成动态规划，也就是自底向上的解决方法，这样可以减少递归调用函数的开销和空间的开销（体现在每层递归函数传的参数，递归层里面申请的变量等）。
