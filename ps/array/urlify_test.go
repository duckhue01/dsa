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

// URLify: write a method to replace all spaces in a string with '%20'. You may
// assume that the string has sufficient space at the end to hold the additional
// characters, and that you ware given the "true" length of the string
// example:
// Input: "Mr John Smith ", 13
// Output: "Mr%20John%20Smith"
func URLify01(str []rune, length int) []rune {
	sps := 0

	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			sps++
		}
	}

	newStr := make([]rune, length+(2*sps))
	x := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			newStr[x] = '%'
			newStr[x+1] = '2'
			newStr[x+2] = '0'
			x += 3
		} else {
			newStr[x] = str[i]
			x++
		}
	}
	return newStr
}

// Wrapping up:
// time complexity: O(n)
// space complexity: O(n)
