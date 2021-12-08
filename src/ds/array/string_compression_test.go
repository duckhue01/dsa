package array

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test case #1", args{"aaabbb"}, "a3b3"},
		{"test case #2", args{"abc"}, "abc"},
		{"test case #3", args{"abba"}, "abba"},
		{"test case #4", args{"abbbba"}, "abbbba"},
		{"test case #5", args{"abbbbbba"}, "a1b6a1"},
		{"test case #6", args{"a"}, "a"},
		{"test case #7", args{"aa"}, "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.str); got != tt.want {
				t.Errorf("compress() = %q, want %q", got, tt.want)
			}
		})
	}
}
