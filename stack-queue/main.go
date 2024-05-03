package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

// push adds an element to the stack
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

// pop removes the last element from the stack
func (s *Stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

type queue struct {
	items []int
}

func (q *queue) enqueue(i int) {
	q.items = append(q.items, i)
}
func (q *queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
func main() {
	myStack := Stack{}
	myQueue := queue{}

	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	myQueue.enqueue(400)
	myQueue.enqueue(500)
	myQueue.enqueue(600)
	myQueue.enqueue(700)
	myQueue.enqueue(800)
	myQueue.enqueue(900)
	myQueue.enqueue(1000)

	myQueue.dequeue()
	myQueue.dequeue()

	fmt.Println(myQueue)

	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	myStack.push(400)
	myStack.push(500)
	myStack.push(600)
	myStack.push(700)
	myStack.push(800)
	myStack.push(900)
	myStack.push(1000)
	myStack.pop()
	myStack.pop()

	fmt.Println(myStack)
}
