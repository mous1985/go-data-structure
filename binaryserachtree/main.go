package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

var count int

func (t *Node) insert(k int) {
	if t.key < k {
		if t.right == nil {
			t.right = &Node{key: k}

		}
		t.right.insert(k)
	} else if t.key > k {
		if t.left == nil {
			t.left = &Node{key: k}

		}
		t.left.insert(k)
	}

}

func (t *Node) search(k int) bool {
	count++
	if t == nil {
		return false
	} else if t.key < k {

		return t.right.search(k)
	} else if t.key > k {

		return t.left.search(k)
	}
	return true
}
func main() {

	tree := &Node{key: 100}
	tree.insert(200)
	tree.insert(50)
	tree.insert(300)
	tree.insert(39)
	tree.insert(59)
	tree.insert(30)

	fmt.Println(tree)
	fmt.Println(tree.search(30))
	fmt.Println(count)
}
