package main

import "fmt"

// alphabet size is the number of possible charachters in tries
const ALPHABET_SIZE = 26

// trie node
type Node struct {
	children [ALPHABET_SIZE]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

// init a new Trie
func initTrie() *Trie {
	return &Trie{root: &Node{}}
}

// search will take a word and search for it in the trie

func (t *Trie) search(word string) bool {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

// insert will take a word and insert it into the trie
func (t *Trie) insert(word string) {
	wordLen := len(word)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// delete will take a word and delete it from the trie
func (t *Trie) delete(word string) {
	t.deleteHelper(t.root, word, 0)
}

func main() {
	fmt.Println("Trie")
}
