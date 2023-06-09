package linked_list

import (
	"errors"
	"fmt"
	"testing"

	"github.com/duckhue01/ps/data/list"
)

// Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.

// 1 -> 2 -> 3 -> 4 : k = 2 => return 3

func ReturnKthToLastWithLen(l *list.S, n int) (int, error) {
	if n > l.Len || n < 1 {
		return 0, errors.New("invalid index")
	}

	c := 1
	temp := l.Head

	for c <= (l.Len - n) {
		c++
		temp = temp.Next
	}

	return temp.Data, nil
}

func ReturnKthToLastTwoPointer(l *list.S, n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid position")
	}

	p1 := l.Head
	p2 := l.Head

	c := 1
	for c < n {
		if p1.Next == nil {
			return 0, errors.New("invalid position")
		}

		p1 = p1.Next
		c++

	}

	for p1.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	return p2.Data, nil
}

func ReturnKthToLastRecursion(l *list.S, n int) (int, error) {
	temp := l.Head

	if n < 1 {
		return 0, errors.New("invalid position")

	}

	var counter int
	res := helper(temp, n, &counter)
	if res == nil {
		return 0, errors.New("invalid position")
	}

	return res.Data, nil
}

func helper(node *list.Node, n int, counter *int) *list.Node {
	if node == nil {
		return nil
	}
	nn := helper(node.Next, n, counter)
	*counter++

	if *counter == n {
		return node
	}

	return nn
}

func TestReturnKthToLast(t *testing.T) {
	type args struct {
		l *list.S
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "return the first to the last element",
			args: args{
				l: makeList([]int{4, 3, 2, 1}),
				n: 1,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "return the second to the last element",
			args: args{
				l: makeList([]int{4, 3, 2, 1}),
				n: 2,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "return an error due to k < 0",
			args: args{
				l: makeList([]int{4, 3, 2, 1}),
				n: 0,
			},
			wantErr: true,
		},
		{
			name: "return an error due to k > l.Len",
			args: args{
				l: makeList([]int{4, 3, 2, 1}),
				n: 100,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("ReturnKthToLastWithLen:", tt.name), func(t *testing.T) {
			got, err := ReturnKthToLastWithLen(tt.args.l, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnKthToLastWithLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReturnKthToLastWithLen() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("ReturnKthToLastTwoPointer:", tt.name), func(t *testing.T) {
			got, err := ReturnKthToLastTwoPointer(tt.args.l, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnKthToLastTwoPointer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReturnKthToLastTwoPointer() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("ReturnKthToLastRecursion:", tt.name), func(t *testing.T) {
			got, err := ReturnKthToLastRecursion(tt.args.l, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnKthToLastRecursion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReturnKthToLastRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
