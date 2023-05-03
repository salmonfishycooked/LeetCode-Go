# [49.Group Anagrams](https://leetcode.cn/problems/group-anagrams)

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

**Example 1:**

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

**Example 2:**

```
Input: strs = [""]
Output: [[""]]
```

**Example 3:**

```
Input: strs = ["a"]
Output: [["a"]]
```

 

**Constraints:**

- `1 <= strs.length <= 104`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.



# 解法：哈希表

**字母异位词** 中每个词所含的字符数量都相同，所以先遍历 `strs` 统计每个 `string` 所含的字符数量，再将其放入哈希表中，最后从哈希表中取值即可。

完整代码如下：

```go
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
```

