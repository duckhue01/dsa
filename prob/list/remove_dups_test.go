package linked_list

import (
	"testing"

	"github.com/duckhue01/ps/data/list"
)

// Remove Dups! Write code to remove duplicates from an unsorted linked list.
// How would you solve this problem if a temporary buffer is not allowed?

func RemoveDupsWithBuffer(l *list.S) {
	t := l.Head

	m := make(map[int]bool)

	var prev *list.Node
	for t != nil {
		if m[t.Data] {
			prev.Next = t.Next
		} else {
			m[t.Data] = true
		}
		prev = t
		t = t.Next

	}
}

func RemoveDups(l *list.S) {
	if l.Len < 2 {
		return
	}
	p1 := l.Head
	for p1 != nil {
		p2 := p1.Next
		prev := p1

		for p2 != nil {
			if p2.Data == p1.Data {
				prev.Next = p2.Next
			}
			prev = p2
			p2 = p2.Next
		}

		p1 = p1.Next
	}
}

func TestRemoveDupsWithBuffer(t *testing.T) {
	type args struct {
		l *list.S
	}
	tests := []struct {
		name string
		args args
		want *list.S
	}{
		{
			name: "#1",
			args: args{
				l: makeList([]int{1, 2, 2, 3, 3, 4, 4}),
			},
			want: makeList([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.l.Show()
			RemoveDupsWithBuffer(tt.args.l)
			tt.args.l.Show()
			// if reflect.DeepEqual(tt.args.l, tt.want) {

			// 	t.Errorf("want %v but got %v", *tt.want, *tt.args.l)
			// }
		})
	}
}

func TestRemoveDups(t *testing.T) {
	type args struct {
		l *list.S
	}
	tests := []struct {
		name string
		args args
		want *list.S
	}{
		{
			name: "#1",
			args: args{
				l: makeList([]int{1, 2, 2, 3, 3, 4, 4}),
			},
			want: makeList([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.l.Show()
			RemoveDups(tt.args.l)
			tt.args.l.Show()
			// if reflect.DeepEqual(tt.args.l, tt.want) {

			// 	t.Errorf("want %v but got %v", *tt.want, *tt.args.l)
			// }
		})
	}
}
