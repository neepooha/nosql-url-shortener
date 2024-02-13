package main

type Node struct {
	data string
	next *Node
}

func Node_Constructor(str string, next *Node) *Node {
	return &Node{
		data: str,
		next: next,
	}
}

type HashNode struct {
	key   string
	value string
	next  *HashNode
}