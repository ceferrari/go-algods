package stack_queue

import "fmt"

type Stack struct {
	top  int
	size int
	arr  []interface{}
}

func NewStack(size int) Stack {
	return Stack{
		top:  -1,
		size: size,
		arr:  make([]interface{}, size),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack) Push(val interface{}) {
	if s.IsFull() {
		fmt.Println("could not push into the stack because it is full")
		return
	}
	s.top++
	s.arr[s.top] = val
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("could not pop from the stack because it is empty")
		return nil
	}
	val := s.arr[s.top]
	s.top--
	return val
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.arr[s.top]
}

func (s *Stack) print() {
	status := "laden"
	if s.IsEmpty() {
		status = "empty"
	}
	if s.IsFull() {
		status = "full"
	}
	fmt.Printf("status: %5s | top: %2d | peek: %+v\n", status, s.top, s.Peek())
}

func RunStack() {
	stack := NewStack(7)

	stack.print()
	stack.Push(1)
	stack.print()
	stack.Push(2)
	stack.print()
	stack.Push(3)
	stack.print()
	stack.Push(4)
	stack.print()
	stack.Push(5)
	stack.print()
	stack.Push(6)
	stack.print()
	stack.Push(7)
	stack.print()
	stack.Push(8)

	fmt.Println()

	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()
	stack.print()
	stack.Pop()

	fmt.Println()

	stack.print()
	stack.Push(1)
	stack.print()
	stack.Pop()
	stack.print()
}
