package linked_list

import (
	"testing"

	"github.com/duckhue01/ps/data/list"
)

func Partition(l *list.S, p int) {
	p1 := l.Head
	p2 := l.Head

	for p2.Next != nil {
		if p1.Data < p && p2.Data < p {
			p1 = p1.Next
		} else if p1.Data >= p && p2.Data < p {
			temp := p1.Data
			p1.Data = p2.Data
			p2.Data = temp
			p1 = p1.Next

		}
		p2 = p2.Next
	}
}

func TestPartition(t *testing.T) {
	l := makeList([]int{5, 4, 6, 1, 2, 8, 10})

	Partition(l, 5)
	l.Show()
}
