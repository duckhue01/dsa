package array

import "testing"

type (
	args1_4 struct {
		str []rune
	}
)

var (
	tests1_4 = []struct {
		name    string
		args1_4 args1_4
		want    bool
	}{
		{
			name: "01",
			args1_4: args1_4{
				str: []rune{'a', 'a'},
			},
			want: true,
		},
		{
			name: "02",
			args1_4: args1_4{
				str: []rune{'a', 'b'},
			},
			want: false,
		},

		{
			name: "03",
			args1_4: args1_4{
				str: []rune{'a'},
			},
			want: true,
		},
		{
			name: "04",
			args1_4: args1_4{
				str: []rune{'a', 'b', 'c'},
			},
			want: false,
		},
	}
)

func Test_palindromePermutation01(t *testing.T) {

	for _, tt := range tests1_4 {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePermutation01(tt.args1_4.str); got != tt.want {
				t.Errorf("palindromePermutation01() = %v, want %v", got, tt.want)
			}
		})
	}
}
