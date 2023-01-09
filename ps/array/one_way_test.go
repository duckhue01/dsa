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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func oneAway(str1 string, str2 string) (bool, error) {
	gap := len(str2) - len(str1)
	// counter := 0
	if abs(gap) > 1 {
		return false, nil
	}

	if gap == 0 {
		c := 0
		for i := 0; i < len(str2); i++ {
			if str1[i] != str2[i] {
				c++
			}
			if c > 1 {
				return false, nil
			}
		}
		return true, nil
	}

	// abs(gap) = 1 or -1
	for i := 0; i < min(len(str1), len(str2)); i++ {
		if str1[i] != str2[i] {
			return false, nil
		}
	}
	return true, nil
}

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
				str1: "abc",
				str2: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#1",
			args: args{
				str1: "abc",
				str2: "abcd",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#1",
			args: args{
				str1: "abc",
				str2: "abdd",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "#1",
			args: args{
				str1: "ebc",
				str2: "abd",
			},
			want:    false,
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

func TestOneWayProperties(t *testing.T) {

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

}
