package main

import (
	"fmt"
)

const (
	maxKeys = 2 // maximum number of keys in a node
)

type node struct {
	keys     [maxKeys + 1]int   // keys in the node
	children [maxKeys + 2]*node // pointers to child nodes
	numKeys  int                // number of keys in the node
}

func (n *node) isLeaf() bool {
	return n.children[0] == nil
}

type btree struct {
	root *node
}

func (t *btree) insert(key int) {
	if t.root == nil {
		t.root = &node{}
		t.root.keys[0] = key
		t.root.numKeys = 1
		return
	}
	split, k := t.root.insert(key)
	if split != nil {
		newRoot := &node{}
		newRoot.keys[0] = k
		newRoot.children[0] = t.root
		newRoot.children[1] = split
		newRoot.numKeys = 1
		t.root = newRoot
	}
}

// 12345
func (n *node) insert(key int) (*node, int) {
	i := n.numKeys - 1
	for i >= 0 && key < n.keys[i] {
		i--
	}
	if n.isLeaf() {
		// Shift keys to make room for new key
		for j := n.numKeys - 1; j > i; j-- {
			n.keys[j+1] = n.keys[j]
		}
		n.keys[i+1] = key
		n.numKeys++
		if n.numKeys > maxKeys {
			return n.split()
		}
		return nil, 0
	}
	split, k := n.children[i+1].insert(key)
	if split != nil {
		// Shift keys to make room for new key
		for j := n.numKeys - 1; j > i; j-- {
			n.keys[j+1] = n.keys[j]
		}
		for j := n.numKeys; j > i+1; j-- {
			n.children[j+1] = n.children[j]
		}
		n.keys[i+1] = k
		n.children[i+2] = split
		n.numKeys++
		if n.numKeys > maxKeys {
			return n.split()
		}
	}
	return nil, 0
}

func (n *node) split() (*node, int) {
	mid := n.numKeys / 2
	newnode := &node{
		numKeys: n.numKeys - mid - 1,
	}
	copy(newnode.keys[:newnode.numKeys], n.keys[mid+1:])
	if !n.isLeaf() {
		copy(newnode.children[:newnode.numKeys+1], n.children[mid+1:])
	}
	n.numKeys = mid
	if n == (&btree{}).root {
		return newnode, n.keys[mid]
	}
	return newnode, n.keys[mid]
}

func (n *node) traverse() {
	fmt.Print("[")
	for i := 0; i < n.numKeys; i++ {
		if !n.isLeaf() {
			n.children[i].traverse()
		}
		fmt.Printf("%d ", n.keys[i])
	}
	if !n.isLeaf() {
		n.children[n.numKeys].traverse()
	}
	fmt.Print("]")
}

func main() {
	t := &btree{}
	t.insert(1)
	t.insert(2)
	t.insert(3)
	t.insert(4)
	t.insert(5)
	t.insert(6)
	t.insert(7)
	t.insert(8)
	t.insert(9)
	t.insert(10)
	t.insert(11)
	t.insert(12)
	t.insert(13)
	t.insert(14)
	t.insert(15)
	t.root.traverse()
}
