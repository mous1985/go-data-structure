package main

import "fmt"

//MaxHeap struct has a slice that holds the array

type MaxHeap struct {
	array []int
}

// insret method adds an element to the heap
func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extarct returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	// when the array is empty

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		// Swap the parent and current node
		h.swap(parent(index), index)
		// Update the index to the parent node
		index = parent(index)
	}
}

//get the parent index of the node

func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index of the node
func leftChild(i int) int {
	return 2*i + 1
}

// swap keys in the array
func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// get the right child index of the node
func rightChild(i int) int {
	return 2*i + 2
}
func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 40}
	for _, v := range buildHeap {
		m.insert(v)
		fmt.Println(m)
	}
}
