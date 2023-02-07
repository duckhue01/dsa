package list

import "fmt"

type sNode struct {
	next *sNode
	data int
}

type SList struct {
	head *sNode
}

func New() *SList {
	return &SList{
		head: nil,
	}
}

func (l *SList) Append(v int) {
	n := &sNode{
		next: nil,
		data: v,
	}
	if l.head != nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
}

func (l *SList) Show() {
	t := l.head
	for t == nil {
		fmt.Printf("%d: ", t.data)
		t = t.next
	}
}
