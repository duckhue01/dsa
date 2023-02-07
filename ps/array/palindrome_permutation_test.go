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

// palindrome permutation: given a string, write a function to check if it
// is a permutation of palindrome. A palindrome is word or phase that is same
// forwards and backwards. A permutation is a rearrangement of letter.
// the palindrome does not need to be limited to just dictionary words.

// EX:
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco eta", etc.)

func palindromePermutation01(str []rune) bool {
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
		count := 0
		for _, v := range dic {
			if v%2 != 0 {
				count++
			}
			if count > 1 {
				return false
			}
		}
		return true
	}
}
