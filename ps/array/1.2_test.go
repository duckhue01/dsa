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

// check permutation: given two strings, write a method to decide if one is
// a permutation of the other.

// idea:
// gradually add each element of add each element of array 1 to map at the same time add it by one
// do the same thing with array 2
// compare length of two map and each element in that map.
func checkPermutation01(str1 []rune, str2 []rune) bool {
	if len(str1) != len(str2) {
		return false
	}
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, v := range str1 {
		map1[v]++
	}
	for _, v := range str2 {
		map2[v]++
	}
	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}

	return true
}

// wrapping up:
// time complexity: O(n)
// space complexity: O(n)

func checkPermutation02(str1 []rune, str2 []rune) bool {
	if len(str1) != len(str2) {
		return false
	}
	temp := [128]int{}

	for _, v := range str1 {
		temp[v]++
	}
	for _, v := range str2 {
		if i := temp[v]; i < 0 {
			return false
		}
	}
	return true
}

// wrapping up:
// time complexity: O(n)
// space complexity: O(n)
