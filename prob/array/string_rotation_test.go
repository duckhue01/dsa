package array

import "testing"

// String Rotation: Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring (e.g.,"waterbottle" is a rotation of"erbottlewat").
//  Hints:#34,#88, #704

func stringRotation(a, b string) bool {
	return isSubstring('a', []rune{'a'})
}

func isSubstring(a rune, b []rune) bool {
	for i := 0; i < len(b); i++ {
		if b[i] == a {
			return true
		}
	}
	return false
}

func Test_stringRotation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringRotation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("stringRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_stringRotation(b *testing.B) {
}

// func Fuzz_stringRotation(f *testing.F) {
// }
