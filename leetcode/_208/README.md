> [LeetCode 208][https://leetcode.cn/problems/implement-trie-prefix-tree]

# Go 实现

```go
type Trie struct {
	head *treeNode
}

type treeNode struct {
	pass, end int
	children  [26]*treeNode
}

func Constructor() Trie {
	return Trie{
		head: &treeNode{},
	}
}

func (t *Trie) Insert(word string) {
	node := t.head
	node.pass++
	for i, n := 0, len(word); i < n; i++ {
		num := word[i] - 'a'
		if node.children[num] == nil {
			newNode := &treeNode{}
			node.children[num] = newNode
		}
		node = node.children[num]
		node.pass++
	}
	node.end++
}

func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	if node == nil || node.end == 0 {
		return false
	}
	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchNode(prefix)
	if node == nil {
		return false
	}
	return true
}

func (t *Trie) searchNode(word string) *treeNode {
	node := t.head
	for i, n := 0, len(word); i < n; i++ {
		num := word[i] - 'a'
		if node.children[num] == nil {
			return nil
		}
		node = node.children[num]
	}
	return node
}
```

> 时间复杂度：$O(|S|)$，$|S|$ 为每次插入或查询的字符串长度。
>
> 空间复杂度：$O(|T| \times Σ)$，$|T|$ 为所有不同字符串的长度之和，$Σ$ 为字符集的大小，在这里 $Σ = 26$。

