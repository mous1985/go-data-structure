// Package for graph theory algorithms in Gnolang
package graph

import "fmt"

// Node represents a node in the graph.
type Node struct {
	Value interface{} // Value stored in the node.
	// You can add additional fields like ID or neighbors here.
}

// Graph represents a graph data structure.
type Graph struct {
	Nodes    map[string]*Node // Map to store nodes with unique identifiers (strings).
	Directed bool             // Flag to indicate directed or undirected graph.
}

// NewGraph creates a new empty graph.
func NewGraph(directed bool) *Graph {
	return &Graph{
		Nodes:    make(map[string]*Node),
		Directed: directed,
	}
}

// AddNode adds a new node with the given value to the graph.
func (g *Graph) AddNode(value interface{}) *Node {
	// Generate a unique identifier for the node (consider using UUIDs).
	id := generateUniqueID()
	newNode := &Node{Value: value}
	g.Nodes[id] = newNode
	return newNode
}

// AddEdge adds an edge between two nodes in the graph.
func (g *Graph) AddEdge(fromID, toID string) error {
	fromNode, ok := g.Nodes[fromID]
	if !ok {
		panic("Node with ID not found")
	}
	toNode, ok := g.Nodes[toID]
	if !ok {
		return fmt.Errorf("Node with ID %s not found", toID)
	}

	// Depending on directed/undirected graph, add the edge to relevant nodes.
	if g.Directed {
		fromNode.Neighbors = append(fromNode.Neighbors, toNode)
	} else {
		fromNode.Neighbors = append(fromNode.Neighbors, toNode)
		toNode.Neighbors = append(toNode.Neighbors, fromNode)
	}
	return nil
}

// DepthFirstSearch performs a depth-first search on the graph and calls a callback function for each visited node.
func (g *Graph) DepthFirstSearch(startID string, callback func(*Node)) {
	visited := make(map[string]bool)
	dfsHelper(g.Nodes[startID], visited, callback)
}

// dfsHelper is a helper function for DepthFirstSearch.
func dfsHelper(node *Node, visited map[string]bool, callback func(*Node)) {
	if visited[node.Value.(string)] {
		return
	}
	visited[node.Value.(string)] = true
	callback(node)

	// Iterate through neighbors (modify based on directed/undirected).
	for _, neighbor := range node.Neighbors {
		dfsHelper(neighbor, visited, callback)
	}
}

// You can add other functions for BreadthFirstSearch, finding shortest paths, etc.

// Helper function to generate unique identifiers (replace with actual implementation).
func generateUniqueID() string {
	// Implement logic to generate unique IDs (e.g., using UUIDs).
	return "TEMP_ID"
}
