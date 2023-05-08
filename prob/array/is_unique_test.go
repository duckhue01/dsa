package array

import (
	"errors"
	"math/big"
	"testing"
)

// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

// confusion:
// + how many distinct characters will be on string?
// + what is the encoding algorithms we use to encode the character?

// idea:
// the reason why i using string instead of string is string is array of byte and go using utf8 to encode char so when a char need more than 8 bit to encode to present one char will need more than one byte what is the encoding algorithm we using utf8 or ascii
// if we cannot using additional data structure we will using two for loop but the TC is O(n^2). Instead, we use the bit presentation and bitwise operator to check and store the existed characters, the TC is O(n)

// conclusion:
// using for range instead of normal for loop because for range will loop over the rune

const (
	ERR_NOT_IN_RANGE = "Char not in valid range"
)

// idea:
// using map to check if a character occurred in map or not. This way we don't have make a whole loop to find duplication. we merely using O(1) operation map produce to check whether or not string is existed.
// time complexity: O(n) in worse case
func isUniqueUsingMap(s string) bool {
	// in case string s merely include ascii characters
	if len(s) > 128 {
		return false
	}

	// map using 8 byte to build struct and 1 byte for each element in map
	charsM := make(map[rune]bool)

	for _, v := range s {
		if !charsM[v] {
			charsM[v] = true
		} else {
			// we return immediately after map already exist current character
			return false
		}
	}

	return true
}

// idea: the minimal accessible unit of memory is one byte so even if we set value type is boolean it will using 8 bit to present to reduce the memory usage we will using 64 bit to present 64 possible character and use & bitwise operation to check if char is presented and [|, ^] to merge bit string
// time complexity: O(n) in worse case
func isUniqueUsingInt64(s string) (bool, error) {

	const (
		MIN_CHAR = 64
		MAX_CHAR = 127
	)

	var (
		acc int64 = 0
	)

	for _, v := range s {
		if v < MIN_CHAR || v > MAX_CHAR {
			return false, errors.New(ERR_NOT_IN_RANGE)
		}
		v -= MIN_CHAR

		curr := int64(1 << v)

		if acc&curr != 0 {
			return false, nil
		}
		// we can use [|, ^] at here because have no two bits have & = 1
		acc = acc ^ curr
	}

	return true, nil

}

// idea:
// time complexity: O(n) in worse case
func isUniqueUsingBigInt(s string) bool {
	acc := big.NewInt(0)

	for _, v := range s {
		temp := big.NewInt(0)
		curr := big.NewInt(1)

		curr.Lsh(curr, uint(v))
		if temp.And(curr, acc).Cmp(big.NewInt(0)) != 0 {
			return false
		}

		acc.Xor(acc, curr)

	}
	return true
}

func Test_isUniqueUsingInt64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "expect ERR_NOT_IN_RANGE when input invalid values",
			args: args{
				s: "123",
			},
			want:    false,
			wantErr: errors.New(ERR_NOT_IN_RANGE),
		},
		{
			name: "expect true when input three different values",
			args: args{
				s: "abc",
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "expect false when input set have same value",
			args: args{
				s: "aab",
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "expect true when input have one value",
			args: args{
				s: "a",
			},
			want:    true,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isUniqueUsingInt64(tt.args.s)
			if (err != nil) != (tt.wantErr != nil) {
				t.Fatalf("isUniqueUsingInt64() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && errors.Is(err, tt.wantErr) {
				t.Fatalf("isUniqueUsingInt64() error message = %v, want error message %v", err, tt.wantErr)

			}
			if got != tt.want {
				t.Errorf("isUniqueUsingInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUniqueUsingMap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "expect true when input three different values",
			args: args{
				s: "abc",
			},
			want: true,
		},
		{
			name: "expect false when input set have same value",
			args: args{
				s: "aab",
			},
			want: false,
		},
		{
			name: "expect true when input have one value",
			args: args{
				s: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUniqueUsingMap(tt.args.s); got != tt.want {
				t.Errorf("isUniqueUsingMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUniqueUsingBigInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "expect true when input three different values",
			args: args{
				s: "abc",
			},
			want: true,
		},
		{
			name: "expect false when input set have same value",
			args: args{
				s: "aab",
			},
			want: false,
		},
		{
			name: "expect true when input have one value",
			args: args{
				s: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUniqueUsingBigInt(tt.args.s); got != tt.want {
				t.Errorf("isUniqueUsingBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Fuzz_isUnique(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		t.Run("isUniqueUsingMap", func(t *testing.T) {
			_ = isUniqueUsingMap(s)
		})

		t.Run("isUniqueUsingInt64", func(t *testing.T) {
			_, _ = isUniqueUsingInt64(s)
		})

		t.Run("isUniqueUsingBigInt", func(t *testing.T) {
			_ = isUniqueUsingBigInt(s)
		})
	})
}

func Benchmark_isUnique(b *testing.B) {
	data := "abc"

	b.Run("isUniqueUsingMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = isUniqueUsingMap(data)
		}
	})

	b.Run("isUniqueUsingInt64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = isUniqueUsingInt64(data)
		}
	})

	b.Run("isUniqueUsingBigInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = isUniqueUsingBigInt(data)
		}
	})
}
