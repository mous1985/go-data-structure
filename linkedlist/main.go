package main

import (
	"fmt"
)

type List struct {
	head   *Node
	length int
}
type Node struct {
	data int
	next *Node
}

var count int

func (l *List) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l List) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}
func (l *List) searchValue(value int) (bool, int) {
	toSearch := l.head
	length := l.length // Add a separate variable to keep track of the remaining length
	for length != 0 {  // Use the separate variable in the loop condition
		if toSearch.data == value {
			return true, count
		}
		toSearch = toSearch.next
		length-- // Decrement the separate variable
		count++
	}
	return false, count
}
func (l *List) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := List{}
	node1 := &Node{data: 45}
	node2 := &Node{data: 90}
	node3 := &Node{data: 30}
	node4 := &Node{data: 60}
	node5 := &Node{data: 75}
	node6 := &Node{data: 15}
	node7 := &Node{data: 105}
	node8 := &Node{data: 120}
	node9 := &Node{data: 135}
	node10 := &Node{data: 150}
	node11 := &Node{data: 165}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.prepend(node8)
	myList.prepend(node9)
	myList.prepend(node10)
	myList.prepend(node11)
	fmt.Println(myList)
	fmt.Println(myList.searchValue(120))
	myList.printData()
	fmt.Println(count)
}
