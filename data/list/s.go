package list

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type S struct {
	Head *Node
	Len  int
}

func (l *S) Append(v int) {
	n := &Node{
		Data: v,
	}
	l.Len += 1

	if l.Head == nil {
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}

func (l *S) Show() {
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
