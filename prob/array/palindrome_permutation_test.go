package array

import "testing"

// palindrome permutation: given a string, write a function to check if it
// is a permutation of palindrome. A palindrome is word or phase that is same
// forwards and backwards. A permutation is a rearrangement of letter.
// the palindrome does not need to be limited to just dictionary words.
// EX:
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco eta", etc.)

// confusing:
// idea:
// assumption:
// conclusion:

func Test_palindromePermutation01(t *testing.T) {

	type args struct {
		str []rune
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "#1",
			args: args{
				str: []rune{'a', 'a'},
			},
			want: true,
		},
		{
			name: "#2",
			args: args{
				str: []rune{'a', 'b'},
			},
			want: false,
		},

		{
			name: "#3",
			args: args{
				str: []rune{'a'},
			},
			want: true,
		},
		{
			name: "#4",
			args: args{
				str: []rune{'a', 'b', 'c'},
			},
			want: false,
		},
		{
			name: "#5 expect true",
			args: args{
				str: []rune{'a', 'b', 'a', 'b'},
			},
			want: true,
		},
		{
			name: "#6 expect false",
			args: args{
				str: []rune{'a', 'b', 'b', 'b'},
			},
			want: false,
		},
		{
			name: "#7 expect true",
			args: args{
				str: []rune{'a', 'b', 'b', 'b', 'a'},
			},
			want: true,
		},
		{
			name: "#8 expect false",
			args: args{
				str: []rune{'a', 'b', 'b', 'c', 'd'},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePermutation(tt.args.str); got != tt.want {
				t.Errorf("palindromePermutation01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func palindromePermutation(str []rune) bool {
	dic := make(map[rune]int)

	for _, v := range str {
		dic[v]++
	}

	if len(str)%2 == 0 {
		for _, v := range dic {
			if v%2 != 0 {
				return false
			}
		}
		return true
	} else {
		opIn := false
		for _, v := range dic {
			if v%2 != 0 {
				if opIn {
					return false
				}
				opIn = true
			}
		}
		return true
	}
}
