package btree

import "fmt"

// implement a in mem btree with item is int8 numbers
// first implement in memory version
// then using it to access disk data

const (
	Order int8 = 5
)

type Btree struct {
	root *node
}

func New() *Btree {
	return &Btree{
		root: new(node),
	}
}

func (b *Btree) Insert(key int8) {

	// In case root doesn't have any element
	if b.root == nil {
		b.root.keys[0] = key
		b.root.len = 1
		return
	}

	split, key := b.root.insert(key)

	if split != nil {
		b.root = &node{
			keys:     [6]int8{key},
			children: [7]*node{b.root, split},
			len:      1,
		}
	}
}

func (n node) Traverse() {
	fmt.Println("hello world")
}
