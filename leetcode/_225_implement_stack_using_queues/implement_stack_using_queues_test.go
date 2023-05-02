package _225_implement_stack_using_queues

import (
	"fmt"
	"testing"
)

func TestProblem225(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
