package main

type Node struct {
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
}
