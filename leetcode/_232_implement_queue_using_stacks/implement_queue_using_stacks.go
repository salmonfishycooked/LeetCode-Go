package _232_implement_queue_using_stacks

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
