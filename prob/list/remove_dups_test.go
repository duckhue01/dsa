package linked_list

import (
	"reflect"
	"testing"

	"github.com/duckhue01/ps/ds/list"
)

// Remove Dups! Write code to remove duplicates from an unsorted linked list.
// How would you solve this problem if a temporary buffer is not allowed?

func removeDupsWithTempBuffer(l *list.SList) *list.SList {

	return list.New()

}

func createListFromSlice(s []int) *list.SList {
	l := list.New()
	for _, v := range s {
		l.Append(v)
	}
	return l
}

func Test_removeDupsWithTempBuffer(t *testing.T) {
	type args struct {
		l *list.SList
	}
	tests := []struct {
		name string
		args args
		want *list.SList
	}{
		{
			name: "",
			args: args{},
			want: &list.SList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDupsWithTempBuffer(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDupsWithTempBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
