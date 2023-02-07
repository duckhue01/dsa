package list_test

import (
	"reflect"
	"testing"

	"github.com/duckhue01/ps/ds/list"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *list.SList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
