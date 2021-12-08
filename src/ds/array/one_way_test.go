package array

import "testing"

func Test_oneWay(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test case #7", args{"paleeee",  "pleee"}, false},
		{"test case #6", args{"pale",  "aaaaaa"}, false},
		{"test case #1", args{"pale",  "ple"}, true},
		{"test case #2", args{"pales",  "pale"}, true},
		{"test case #3", args{"pale",  "bale"}, true},
		{"test case #4", args{"pale",  "bake"}, false},
		{"test case #5", args{"pale",  "elap"}, false},
		{"test case #8", args{"",  ""}, true},
		{"test case #9", args{"duckhue01",  "duckhue02"}, true},
		{"test case #10", args{"duckhue01",  "duckhue01"}, true},
		{"test case #11", args{"duckhue01",  ""}, false},
		{"test case #12", args{"",  "duckhue02"}, false},
		{"test case #13", args{"aaa",  "bbb"}, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneWay(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("oneWay() = %v, want %v", got, tt.want)
			}
		})
	}
}
