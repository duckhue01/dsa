package array

import "testing"

type (
	args1_2 struct {
		str1 []rune
		str2 []rune
	}
)

var (
	tests1_2 = []struct {
		name    string
		args1_2 args1_2
		want    bool
	}{
		{
			name: "01",
			args1_2: args1_2{
				str1: []rune{'a'},
				str2: []rune{'a'},
			},
			want: true,
		},
		{
			name: "02",
			args1_2: args1_2{
				str1: []rune{'a', 'b'},
				str2: []rune{'a'},
			},
			want: false,
		},

		{
			name: "01",
			args1_2: args1_2{
				str1: []rune{'a', 'b'},
				str2: []rune{'b', 'a'},
			},
			want: true,
		},
		{
			name: "01",
			args1_2: args1_2{
				str1: []rune{'a', 'b', 'c'},
				str2: []rune{'c', 'a', 'b'},
			},
			want: true,
		},
	}
)

func Test_checkPermutation01(t *testing.T) {
	for _, tt := range tests1_2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPermutation01(tt.args1_2.str1, tt.args1_2.str2); got != tt.want {
				t.Errorf("checkPermutation01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPermutation02(t *testing.T) {
	for _, tt := range tests1_2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPermutation02(tt.args1_2.str1, tt.args1_2.str2); got != tt.want {
				t.Errorf("checkPermutation01() = %v, want %v", got, tt.want)
			}
		})
	}
}
