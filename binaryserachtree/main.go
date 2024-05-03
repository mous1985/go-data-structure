package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

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

func search() {}
func main() {

	tree := &Node{key: 100}
	tree.insert(50)

	fmt.Println(tree)

}
