package stacks

import "fmt"

type stack[T any] struct {
	Length  func() int
	IsEmpty func() bool
	Push    func(val T)
	Pop     func() T
	Peek    func() T
	Print   func()
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	len := func() int {
		return len(slice)
	}
	top := func() int {
		return len() - 1
	}
	return stack[T]{
		Length: func() int {
			return len()
		},
		IsEmpty: func() bool {
			return len() == 0
		},
		Push: func(val T) {
			slice = append(slice, val)
		},
		Pop: func() T {
			if len() == 0 {
				return *new(T)
			}
			val := slice[top()]
			slice = slice[:top()]
			return val
		},
		Peek: func() T {
			if len() == 0 {
				return *new(T)
			}
			return slice[top()]
		},
		Print: func() {
			status := "empty"
			peek := *new(T)
			if len() > 0 {
				peek = slice[top()]
				status = "laden"
			}
			fmt.Printf("status: %5s | len: %2d | peek: %v\n", status, len(), peek)
		},
	}
}

func RunStack() {
	stack := Stack[int]()
	stack.Print()
	stack.Push(1)
	stack.Print()
	stack.Push(2)
	stack.Print()
	stack.Push(3)
	stack.Print()
	stack.Push(4)
	stack.Print()
	stack.Push(5)
	stack.Print()
	stack.Push(6)
	stack.Print()
	stack.Push(7)
	stack.Print()
	stack.Push(8)
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Push(1)
	stack.Print()
	stack.Pop()
	stack.Print()
}
