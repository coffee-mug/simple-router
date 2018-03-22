package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	root := &Node{value: 0, children: [256]*Node{}}

	node := root.Add("lol", 0, 10)

	assert.Equal(t, node.value, 10, "n.Add(value) should return a node with the same value")
}

func TestGet(t *testing.T) {
	root := &Node{value: 0, children: [256]*Node{}}

	_ = root.Add("Fantastic", 0, 10)
	_ = root.Add("Georges", 0, 1800)
	_ = root.Add("Poney", 0, 54)
	_ = root.Add("lol", 0, 10)

	got, _ := root.Traverse("Poney", 0)
	assert.Equal(t, got, 54, "n.Traverse(key) should return the value at that key")

	got, _ = root.Traverse("Georges", 0)
	assert.Equal(t, got, 1800, "n.Traverse(key) should return the value at that key")

	got, _ = root.Traverse("lol", 0)
	assert.Equal(t, got, 10, "n.Traverse(key) should return the value at that key")

	got, _ = root.Traverse("Fantastic", 0)
	assert.Equal(t, got, 10, "n.Traverse(key) should return the value at that key")
}
