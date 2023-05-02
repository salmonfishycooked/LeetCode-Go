package _232_implement_queue_using_stacks

import (
	"fmt"
	"testing"
)

func TestProblem232(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	q.Push(5)
	for !q.Empty() {
		fmt.Println(q.Pop())
	}
}
