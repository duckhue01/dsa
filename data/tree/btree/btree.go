package btree

// implement a in mem btree with item is int8 numbers
// first implement in memory version
// then using it to access disk data

// For each node x, the keys are stored in increasing order.
// In each node, there is a boolean value x.leaf which is true if x is a leaf.
// If n is the order of the tree, each internal node can contain at most n - 1 keys along with a pointer to each child.
// Each node except root can have at most n children and at least n/2 children.
// All leaves have the same depth (i.e. height-h of the tree).
// The root has at least 2 children and contains a minimum of 1 key.
// If n ≥ 1, then for any n-key B-tree of height h and minimum degree t ≥ 2, h ≥ logt (n+1)/2.
