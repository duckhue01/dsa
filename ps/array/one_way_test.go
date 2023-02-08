package array

import (
	"testing"
)

// one away: there are three types of edits that can be performed on string
// insert a character, remove a character, or replace a character. Given
// two strings, write a function to check if they are one edit(or zero edit) away

// EXAMPLE
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false
// Hints:#23, #97, #130

// confusion:

// idea:

// conclusion:

// todo: using string and for range loop

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func oneAway(str1 string, str2 string) (bool, error) {
	r1 := []rune(str1)
	r2 := []rune(str2)

	gap := len(r2) - len(r1)
	if abs(gap) > 1 {
		return false, nil
	}

	switch gap {
	case 0:
		c := 0
		for i := range r1 {
			if r1[i] != r2[i] {
				c++
			}
			if c > 1 {
				return false, nil
			}

		}

		return true, nil

	case -1:
		// where we have remove operation
		i, j, c := 0, 0, 0

		for j < len(r2) {
			if r1[i] != r2[j] {
				c++
				i++
			} else {
				i++
				j++
			}
			if c > 1 || (c == 1 && j == len(r2)-1 && i < len(r1)-1) {
				return false, nil
			}

		}
		return true, nil

	case 1:
		// where we have insert operation
		i, j, c := 0, 0, 0

		for i < len(r1) {
			if r1[i] != r2[j] {
				c++
				j++
			} else {
				i++
				j++
			}
			if c > 1 || (c == 1 && i == len(r1)-1 && j < len(r2)-1) {
				return false, nil
			}

		}
		return true, nil

	}

	return true, nil
}

// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false

func TestOneAway(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				str1: "pale",
				str2: "ple",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#2",
			args: args{
				str1: "pales",
				str2: "pale",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#3",
			args: args{
				str1: "pale",
				str2: "bale",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#4",
			args: args{
				str1: "pale",
				str2: "bake",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "#5",
			args: args{
				str1: "abc",
				str2: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#6",
			args: args{
				str1: "abc",
				str2: "abcd",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#7",
			args: args{
				str1: "abc",
				str2: "abdd",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "#8",
			args: args{
				str1: "ebc",
				str2: "abd",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "str1 have more one character than str2",
			args: args{
				str1: "abcd",
				str2: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "str2 have more one character than str1",
			args: args{
				str1: "abc",
				str2: "abcd",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := oneAway(tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("OneAway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzOneWay(f *testing.F) {
	f.Fuzz(func(t *testing.T, s1, s2 string) {
		res, err := oneAway(s1, s2)
		if err != nil {
			t.Error(err)
		}
		t.Log(res)
	})
}

func BenchmarkOneWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := oneAway("abc", "abcd")
		if err != nil {
			b.Error(err)
		}
	}

}
