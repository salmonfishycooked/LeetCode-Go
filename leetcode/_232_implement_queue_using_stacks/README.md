# [232. Implement Queue using Stacks](https://leetcode.cn/problems/implement-queue-using-stacks/description/)

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `peek`, `pop`, and `empty`).

Implement the `MyQueue` class:

- `void push(int x)` Pushes element x to the back of the queue.
- `int pop()` Removes the element from the front of the queue and returns it.
- `int peek()` Returns the element at the front of the queue.
- `boolean empty()` Returns `true` if the queue is empty, `false` otherwise.

**Notes:**

- You must use **only** standard operations of a stack, which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.
- Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

 

**Example 1:**

```
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
```

 

**Constraints:**

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `peek`, and `empty`.
- All the calls to `pop` and `peek` are valid.

 

**Follow-up:** Can you implement the queue such that each operation is **[amortized](https://en.wikipedia.org/wiki/Amortized_analysis)** `O(1)` time complexity? In other words, performing `n` operations will take overall `O(n)` time even if one of those operations may take longer.



# 解法：双栈

利用一个 `in` 栈，一个 `out` 栈。

有新元素进来的时候 `push` 到 `in` 栈，当需要 `peek` 或 `pop` 的时候，需要把 `in` 中的元素倒入 `out` 中， 此时 `out` 的栈顶元素就是队首元素（当然，如果 `out` 栈非空的话就不用把 `in` 栈元素倒入了，不然会让 `out` 的栈顶元素并非队首元素）

完整代码如下

```go
type MyQueue struct {
	in  *stack
	out *stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  &stack{},
		out: &stack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.in.push(x)
}

func (q *MyQueue) Pop() int {
	q.toOut()
	return q.out.pop()
}

func (q *MyQueue) Peek() int {
	q.toOut()
	return q.out.peek()
}

func (q *MyQueue) Empty() bool {
	return q.in.isEmpty() && q.out.isEmpty()
}

func (q *MyQueue) toOut() {
	if q.out.isEmpty() {
		for !q.in.isEmpty() {
			q.out.push(q.in.pop())
		}
	}
}

type stack struct {
	e []int
}

func (s *stack) push(n int) {
	s.e = append(s.e, n)
}

func (s *stack) peek() int {
	if !s.isEmpty() {
		return s.e[len(s.e)-1]
	}
	return -1
}

func (s *stack) pop() int {
	n := s.peek()
	if !s.isEmpty() {
		s.e = s.e[:len(s.e)-1]
	}
	return n
}

func (s *stack) size() int {
	return len(s.e)
}

func (s *stack) isEmpty() bool {
	return s.size() == 0
}
```

