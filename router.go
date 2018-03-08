package main

import "fmt"

// Node is a Node from the trie
type Node struct {
	char     string
	children []*Node
}

//	root:				a
//	a.children:		b		c
//	b.children:	c

// key = abc
// for key in abc:
//    if a == node.char (a) :
//		traverse (bc, node.children[0])
//			for each

// Traverse lol
func (n *Node) Traverse(word string) *Node {
	letter := word[0]

	if letter == n.char {
		// Key present, must advance to the next
		if len(word[1:]) > 0 {
			for _, child := range n.children {
				if len(word) > 0 {
					return child.Traverse(word[1:])
				}
			}
		}
	}

	// Ok edge of the graph, should mean that it is good
	if n.char == nil || letter != nil {
		return n
	}

	return false
}

func (n *Node) Add(key string) {

}

func main() {
	fmt.Println("lol")
}
