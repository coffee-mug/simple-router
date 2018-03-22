package main

import (
	"fmt"
)

type Node struct {
<<<<<<< HEAD
	value    int
	children [256]*Node
}

//Traverse, recursively get the value at provided key, null otherwise the trie
func (n *Node) Traverse(word string, currentElement int) (int, *Node) {
	if n == nil {
		// We reached the end of the trie
		return 0, n
	}
	if currentElement == len(word) {
		// We reached the end of the trie for that key
		return n.value, n
	}
	return n.children[word[currentElement]].Traverse(word, currentElement+1)
}

func (n *Node) Add(key string, currentElement int, value int) *Node {
	if n == nil {
		return n
	}

	if len(key) == currentElement {
		n.value = value
		return n
	}

	// Add a node at currentElement index
	deeper := &Node{value: 0, children: [256]*Node{}}

	n.children[key[currentElement]] = deeper

	return deeper.Add(key, currentElement+1, value)
=======
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
>>>>>>> ae6220ea4ac2e1807f7c099ab4ecaafacb8d1abd
}

func main() {
	root := &Node{value: 0, children: [256]*Node{}}

	lode := root.Add("lol", 0, 12)

	fmt.Println(lode.value)

	got, _ := root.Traverse("lol", 0)

	fmt.Println(got, 12)
}
