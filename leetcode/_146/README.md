> [LeetCode 146][https://leetcode.cn/problems/lru-cache]

# 双向链表 + 哈希表

```go
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	cache      map[int]*linkedNode
	head, tail *linkedNode
	size       int
	cap        int
}

type linkedNode struct {
	prev, next *linkedNode
	key, data  int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: make(map[int]*linkedNode),
		head:  &linkedNode{},
		tail:  &linkedNode{},
		cap:   capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	l.setNodeLatest(node)
	return node.data
}

func (l *LRUCache) Put(key int, value int) {
	// 该页是否存在
	if l.exist(key, value) {
		return
	}
	// 判断缓存是否已满
	if l.size == l.cap {
		l.handleFull()
	}

	// 放入新页
	node := &linkedNode{
		key:  key,
		data: value,
	}
	node.prev = l.head
	node.next = l.head.next
	node.next.prev = node
	l.head.next = node
	l.cache[key] = node
	l.size++
}

// exist 判断该页是否已经存在，存在则改数据
func (l *LRUCache) exist(key int, value int) bool {
	if node, ok := l.cache[key]; ok {
		node.data = value
		// 将该页放入首部，因为它最近被访问过
		l.setNodeLatest(node)
		return true
	}
	return false
}

// setNodeLatest 将结点放置到首部
func (l *LRUCache) setNodeLatest(node *linkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	node.prev = l.head
}

// handleFull 移除链表尾部元素（该元素即为最近最久未使用的页）
func (l *LRUCache) handleFull() {
	node := l.tail.prev
	l.tail.prev = node.prev
	node.prev.next = l.tail
	delete(l.cache, node.key)
	l.size--
}
```

> 时间复杂度：对于 `Put` 和 `Get` 操作都是 $O(1)$。
>
> 空间复杂度：$O(capacity)$，`capacity` 为 $LRUCache$ 的容量。