package _208

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
