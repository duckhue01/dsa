package linked_list

import (
	"errors"
	"testing"

	"github.com/duckhue01/ps/data/list"
	"github.com/google/go-cmp/cmp"
)

// Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.

func DeleteMiddleNode(l *list.S, p *list.Node) error {
	if p.Next == nil || p == l.Head {
		return errors.New("can not remove the first & last item")
	}
	p.Data = p.Next.Data
	p.Next = p.Next.Next
	return nil
}

func TestDeleteMiddleNode(t *testing.T) {

	t.Run("should remove the middle node", func(t *testing.T) {
		l := makeList([]int{5, 1, 3})

		p := l.Head.Next
		_ = DeleteMiddleNode(l, p)

		if !cmp.Equal(l.ToSlice(), []int{3, 5}) {
			t.Error("not equal")
		}

	})

	t.Run("should not remove the first item", func(t *testing.T) {
		l := makeList([]int{5, 1})

		p := l.Head
		_ = DeleteMiddleNode(l, p)

		if !cmp.Equal(l.ToSlice(), []int{1, 5}) {
			t.Error("not equal")
		}

	})

	t.Run("should not remove the last item", func(t *testing.T) {
		l := makeList([]int{5, 1})

		p := l.Head
		_ = DeleteMiddleNode(l, p)

		if !cmp.Equal(l.ToSlice(), []int{1, 5}) {
			t.Error("not equal")
		}

	})

}
