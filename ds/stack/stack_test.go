package stack

import "testing"

type Node struct {
	val rune
	nxt *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) push() {

}

func (s *Stack) pop() {

}

func TestStack(t *testing.T) {
	a := []int{1, 2, 3}
	t.Log(a)

}
