package btree

type node struct {
	keys     [Order + 1]int8
	children [Order + 2]*node
	len      int8
}

func (n *node) isLeaf() bool {
	return n.children[0] == nil
}

func (n node) insert(key int8) (*node, int8) {
	i := n.len - 1
	for i >= 0 && key < n.keys[i] {
		i--
	}

	if n.isLeaf() {
		for j := n.len - 1; j > i; j-- {
			n.keys[j+1] = n.keys[j]
		}
		n.keys[i+1] = key
		n.len++
		if n.len > Order {
			return n.split()
		}
		return nil, 0
	}
	split, key := n.children[i+1].insert(key)

	if split != nil {
		// Shift keys to make room for new key
		for j := n.len - 1; j > i; j-- {
			n.keys[j+1] = n.keys[j]
		}
		for j := n.len; j > i+1; j-- {
			n.children[j+1] = n.children[j]
		}
		n.keys[i+1] = key
		n.children[i+2] = split
		n.len++
		if n.len > Order {
			return n.split()
		}
	}
	return nil, 0
}

func (n node) split() (*node, int8) {
	mi := n.len / 2
	newNode := &node{
		len: n.len - mi - 1,
	}
	copy(newNode.keys[:newNode.len], n.keys[mi-1:])
	if !n.isLeaf() {
		copy(newNode.children[:newNode.len+1], n.children[mi+1:])
	}
	n.len = mi

	return newNode, n.keys[mi]
}
