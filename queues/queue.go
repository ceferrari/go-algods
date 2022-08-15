package queues

import "fmt"

type queue[T any] struct {
	Length  func() int
	IsEmpty func() bool
	Enqueue func(val T)
	Dequeue func() T
	Peek    func() T
	Print   func()
}

func Queue[T any]() queue[T] {
	slice := make([]T, 0)
	len := func() int {
		return len(slice)
	}
	return queue[T]{
		Length: func() int {
			return len()
		},
		IsEmpty: func() bool {
			return len() == 0
		},
		Enqueue: func(val T) {
			slice = append(slice, val)
		},
		Dequeue: func() T {
			if len() == 0 {
				return *new(T)
			}
			val := slice[0]
			slice = slice[1:]
			return val
		},
		Peek: func() T {
			if len() == 0 {
				return *new(T)
			}
			return slice[0]
		},
		Print: func() {
			status := "empty"
			peek := *new(T)
			if len() > 0 {
				peek = slice[0]
				status = "laden"
			}
			fmt.Printf("status: %5s | len: %2d | peek: %v\n", status, len(), peek)
		},
	}
}

func RunQueue() {
	queue := Queue[int]()
	queue.Print()
	queue.Enqueue(1)
	queue.Print()
	queue.Enqueue(2)
	queue.Print()
	queue.Enqueue(3)
	queue.Print()
	queue.Enqueue(4)
	queue.Print()
	queue.Enqueue(5)
	queue.Print()
	queue.Enqueue(6)
	queue.Print()
	queue.Enqueue(7)
	queue.Print()
	queue.Enqueue(8)
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Enqueue(1)
	queue.Print()
	queue.Dequeue()
	queue.Print()
}
