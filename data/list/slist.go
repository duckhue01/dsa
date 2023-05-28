package list

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type SList struct {
	Head *Node
}

func (l *SList) Append(v int) {
	n := &Node{
		Data: v,
	}

	if l.Head == nil {
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}

func (l *SList) Show() {
	t := l.Head
	for t != nil {
		if t.Next != nil {
			fmt.Printf("%d -> ", t.Data)

		} else {
			fmt.Printf("%d\n", t.Data)
		}
		t = t.Next
	}
}
