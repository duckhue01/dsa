package array

import (
	"testing"
)

// Check Permutation: Given two strings,write a method to decide if one is a permutation of the other.

// confusion:
// is string include unicode character or merely ASCII characters.
// idea:
// put each character into a map and then compare every single character
// if string include unicode character we have to convert it into []rune or using for range to get full unicode element
// assumption:
// string just include ASCII characters
// conclusion:

// idea: we push all of two strings into 2 map respectively then compare these two
// time complexity: O(n) in worse case
// space complexity: O(n) in worse case
func detectPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range a {
		m1[v]++
	}
	for _, v := range b {
		m2[v]++
	}
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

func Test_detectPermutation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "input b is permutation of a, expect true",
			args: args{
				a: "abcd",
				b: "adbc",
			},
			want: true,
		},
		{
			name: "input b isn't permutation of a, expect false",
			args: args{
				a: "abcd",
				b: "aabc",
			},
			want: false,
		},
		{
			name: "input b has len different from a, expect false",
			args: args{
				a: "abcd",
				b: "aa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectPermutation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("detectPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
