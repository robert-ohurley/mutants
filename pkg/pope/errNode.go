package main

import "fmt"

type errorNode struct {
	//The error associated with this error node
	err *err

	//The pointers to the connected error nodes.
	next []*errorNode
}

func (e *errorNode) Error() {
	fmt.Printf("type: %v, expected: %s, got: %s, len of connected nodes = %d\n", e.err.errType, e.err.expected, e.err.recieved, len(e.next))
}

func (e *errorNode) IsLeaf() bool {
	return len(e.next) == 0
}

func (e *errorNode) AddNext(node *errorNode) {
	e.next = append(e.next, node)
}

func NewErrorNode(e *err) *errorNode {
	return &errorNode{e, make([]*errorNode, 0)}
}
