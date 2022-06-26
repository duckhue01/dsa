package array

import "testing"

type (
	args1_1 struct {
		str []rune
	}
)

var (
	tests1_1 = []struct {
		name    string
		args1_1 args1_1
		want    bool
	}{
		{
			name: "01",
			args1_1: args1_1{
				str: []rune{'1', 'a', '1'},
			},
			want: false,
		},
		{
			name: "02",
			args1_1: args1_1{
				str: []rune{'1', 'a', 'a'},
			},
			want: false,
		},
		{
			name: "03",
			args1_1: args1_1{
				str: []rune{'1', 'a', '2'},
			},
			want: true,
		},

		{
			name: "04",
			args1_1: args1_1{
				str: []rune{'1'},
			},
			want: true,
		},
		{
			name: "05",
			args1_1: args1_1{
				str: []rune{'1', '1'},
			},
			want: false,
		},
		{
			name: "06",
			args1_1: args1_1{
				str: []rune{'a', 'a', '1'},
			},
			want: false,
		},
		{
			name: "07",
			args1_1: args1_1{
				str: []rune{'a', 'a', '1'},
			},
			want: false,
		},
		{
			name: "08",
			args1_1: args1_1{
				str: []rune{'a', 'b', 'c'},
			},
			want: true,
		},
	}
)

func Test_isUnique01(t *testing.T) {
	for _, tt := range tests1_1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique01(tt.args1_1.str); got != tt.want {
				t.Errorf("isUnique01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUnique02(t *testing.T) {
	for _, tt := range tests1_1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique02(tt.args1_1.str); got != tt.want {
				t.Errorf("isUnique02() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUnique03(t *testing.T) {

	tests := []struct {
		name    string
		args1_1 args1_1
		want    bool
	}{
		{
			name: "01",
			args1_1: args1_1{
				str: []rune{'b', 'a', 'b'},
			},
			want: false,
		},
		{
			name: "02",
			args1_1: args1_1{
				str: []rune{'b', 'a', 'a'},
			},
			want: false,
		},
		{
			name: "03",
			args1_1: args1_1{
				str: []rune{'b', 'a', 'c'},
			},
			want: true,
		},

		{
			name: "04",
			args1_1: args1_1{
				str: []rune{'a'},
			},
			want: true,
		},
		{
			name: "05",
			args1_1: args1_1{
				str: []rune{'a', 'a'},
			},
			want: false,
		},
		{
			name: "06",
			args1_1: args1_1{
				str: []rune{'a', 'a', 'b'},
			},
			want: false,
		},
		{
			name: "07",
			args1_1: args1_1{
				str: []rune{'a', 'a', 'b'},
			},
			want: false,
		},
		{
			name: "08",
			args1_1: args1_1{
				str: []rune{'a', 'b', 'c'},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique03(tt.args1_1.str); got != tt.want {
				t.Errorf("isUnique03() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_isUnique04(t *testing.T) {
	for _, tt := range tests1_1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique04(tt.args1_1.str); got != tt.want {
				t.Errorf("isUnique02() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_isUnique05(t *testing.T) {
	for _, tt := range tests1_1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique05(tt.args1_1.str); got != tt.want {
				t.Errorf("isUnique02() = %v, want %v", got, tt.want)
			}
		})
	}
}