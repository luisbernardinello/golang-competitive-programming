package main

import "fmt"

type Stack struct {
	items []int
}

// Stack Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Stack Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// Queue Enqueue
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Queue Dequeue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	// Stack operations
	myStack := Stack{}
	fmt.Println("Initial Stack:", myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println("Stack after pushes:", myStack)

	myStack.Pop()
	fmt.Println("Stack after pop:", myStack)

	// Queue operations
	myQueue := Queue{}
	fmt.Println("Initial Queue:", myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println("Queue after enqueues:", myQueue)

	myQueue.Dequeue()
	fmt.Println("Queue after dequeue:", myQueue)
}
