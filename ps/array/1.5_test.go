package array

import "testing"



// one away: there are three types of edits that can be performed on string
// insert a character, remove a character, or replace a character. Given
// two strings, write a function to check if they are one edit(or zero edit) away

// EXAMPLE
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false

func OneAway(str1 []rune, str2 []rune) bool {

	return false
}


func TestOneAway(t *testing.T) {
	type args struct {
		str1 []rune
		str2 []rune
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
			if got := OneAway(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("OneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
