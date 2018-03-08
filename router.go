package main

import "fmt"

// Node is a Node from the trie
type Node struct {
	char     string
  value    int
	children []*Node
}

//  root:                     nil 
//	root.children:				    a
//	a.children:		          b		c
//	b.children:	          c

/*
  for child in children:
    if currentComponent.child.char == letter

    if currentComponent.child.char == nil
*/

func (n *Node) traverse(key string) (bool, *Node) {
  letter := key[0]

  for _, child := range n.children {
    if child.char == letter {
      // start again, with a letter advanced and the child
      if len(key[1:]) > 0 {
        return child.traverse(key[1:])
      }
    }

    if child == nil || child.char == nil {
      // We reached the end of the trie, return false and node reference
      return true, child
    }
  }
}

func (n *Node) Add(n *Node, key string, value int) {
  // first traverse to get the node
  _, node := n.traverse(key)
}

func NewTrie() *Node {
  return &Node{
    char: nil,
    value: 0,
    children: make([]*Node)
  }
}

func main() {
	fmt.Println("lol")
}
