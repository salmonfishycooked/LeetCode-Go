package _225_implement_stack_using_queues

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
