package array

import (
	"reflect"
	"testing"
)

type (
	args1_3 struct {
		str    []rune
		length int
	}
)

var (
	tests1_3 = []struct {
		name    string
		args1_3 args1_3
		want    []rune
	}{
		{
			name: "1",
			args1_3: args1_3{
				str:    []rune{'1', ' ', '2', ' ', ' '},
				length: 3,
			},
			want: []rune{'1', '%', '2', '0', '2'},
		},
	}
)

func TestURLify01(t *testing.T) {

	for _, tt := range tests1_3 {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify01(tt.args1_3.str, tt.args1_3.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLify01() = %v, want %v", got, tt.want)
			}
		})
	}
}
