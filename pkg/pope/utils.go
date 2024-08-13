package main

import (
	"fmt"
	"strings"
)

func PrintTree(startingNode *errorNode) {
	printTree(startingNode, 0)
}

func printTree(node *errorNode, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("\t")
	}

	node.Error()

	if len(node.next) == 0 {
		return
	}

	for _, n := range node.next {
		printTree(n, indent+1)
	}
}

func createParamString(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

func assert(expected, recieved any, desc string) {
	if expected != recieved {
		panic(desc)
	}
}

func assertGreaterThan(expected, recieved int, desc string) {
	if expected < recieved {
		panic(desc)
	}
}
