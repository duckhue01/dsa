package btree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Btree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBtree_Insert(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		key int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Btree{
				root: tt.fields.root,
			}
			b.Insert(tt.args.key)
		})
	}
}

func Test_node_Traverse(t *testing.T) {
	type fields struct {
		keys     [Order + 1]int8
		children [Order + 2]*node
		len      int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				keys:     tt.fields.keys,
				children: tt.fields.children,
				len:      tt.fields.len,
			}
			n.Traverse()
		})
	}
}
