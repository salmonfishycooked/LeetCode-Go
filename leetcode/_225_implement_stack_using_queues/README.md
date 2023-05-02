# [225. Implement Stack using Queues](https://leetcode.cn/problems/implement-stack-using-queues)

Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (`push`, `top`, `pop`, and `empty`).

Implement the `MyStack` class:

- `void push(int x)` Pushes element x to the top of the stack.
- `int pop()` Removes the element on the top of the stack and returns it.
- `int top()` Returns the element on the top of the stack.
- `boolean empty()` Returns `true` if the stack is empty, `false` otherwise.

**Notes:**

- You must use **only** standard operations of a queue, which means that only `push to back`, `peek/pop from front`, `size` and `is empty` operations are valid.
- Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.

 

**Example 1:**

```
Input
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 2, 2, false]

Explanation
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // return 2
myStack.pop(); // return 2
myStack.empty(); // return False
```

 

**Constraints:**

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `top`, and `empty`.
- All the calls to `pop` and `top` are valid.

 

**Follow-up:** Can you implement the stack using only one queue?



# 解法一：双队列

用两个队列 `q` 和 `help` ，当需要 `top` 和 `pop` 操作的时候，将 `q` 的元素倒入 `help` 中，当然 `q` 只剩最后一个的时候要记录它的值用于返回，这个值就是栈顶元素。

```go
type MyStack struct {
	q []int
	help []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.q = append(s.q, x) // 入队
}

func (s *MyStack) Pop() (v int) {
	v = s.Top()
	s.q = s.q[:len(s.q)-1]
	return
}

func (s *MyStack) Top() (v int) {
	for i := len(s.q); i > 0; i-- {
		s.help = append(s.help, s.q[0]) // help队列获取q队列队首元素再入队
		if i == 1 {
			v = s.q[0]
		}
		s.q = s.q[1:] // q队列出队
	}
	s.q, s.help = s.help, s.q // 交换队列的性质
	return
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}
```



# 解法二：单队列

```go
type MyStack struct {
	e []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.e = append(s.e, x) // 入队
}

func (s *MyStack) Pop() (v int) {
	v = s.Top()
	s.e = s.e[:len(s.e)-1]
	return
}

func (s *MyStack) Top() (v int) {
	// 此处是模拟队列弹出元素再插入队尾，直到找到原先队尾的元素返回它，再将其插入队尾
	// 以下代码等价于 return s.e[len(s.e) - 1]
	for i := len(s.e); i > 0; i-- {
		if i == 1 {
			v = s.e[0]
		}
		s.e = append(s.e, s.e[0])
		s.e = s.e[1:]
	}
	return
}

func (s *MyStack) Empty() bool {
	return len(s.e) == 0
}
```

