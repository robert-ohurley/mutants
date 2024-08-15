package main

import "fmt"

type Visitor struct {
	//Where to begin the traversal from.
	root *errorNode

	//This node is the one we are currently visiting which will be exposed to the user.
	currNode *errorNode

	//DFS traversal is done prior to exposing current node to the user.
	stack []*errorNode
}

func NewVisitor(root *errorNode) *Visitor {
	v := Visitor{
		root:     root,
		currNode: nil,
		stack:    make([]*errorNode, 0),
	}

	v.buildStack(v.root)
	return &v

}

func (v *Visitor) CurrNode() *errorNode {
	return v.currNode
}

// Returns true provided that there are more nodes to traverse.
func (v *Visitor) Walk() bool {
	return v.walk()
}

func (v *Visitor) walk() bool {
	if len(v.stack) == 0 {
		return false
	}

	v.currNode = v.stack[0]
	v.stack = v.stack[1:]
	return true
}

func (v *Visitor) buildStack(node *errorNode) {
	if node == nil {
		return
	}

	v.stack = append(v.stack, node)

	for _, n := range node.next {
		v.stack = append(v.stack, n)
		v.buildStack(n)
	}
}

var leafNodes = 0

func PrintTree(startingNode *errorNode) {
	printTree(startingNode, 0)
	fmt.Println("total possible outcomes: ", leafNodes)
}

func printTree(node *errorNode, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("\t")
	}

	if node.IsLeaf() {
		leafNodes++
	}

	node.Error()

	if len(node.next) == 0 {
		return
	}

	for _, n := range node.next {
		printTree(n, indent+1)
	}
}
