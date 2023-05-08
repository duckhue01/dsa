package array

import (
	"fmt"
	"testing"
)

// string compression: implement a method to perform basic string compression
// using the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3,
// if the "compressed" string would not become smaller than the original string
// your method should return the original string. You can assume the string as only uppercase
// and lowercase letters(a-z)

func com(token byte, n int) string {
	if n == 1 {
		return fmt.Sprint(string(token), n)
	}

	return fmt.Sprint(string(token), n)
}

func compression(o string) string {
	if len(o) == 0 {
		return o
	}

	prev := o[0]
	counter := 1
	compressed := ""

	for i := 1; i < len(o); i++ {
		if prev != o[i] {
			compressed += com(prev, counter)
			if len(compressed) >= len(o) {
				return o
			}
			prev = o[i]
			counter = 1
		} else {
			counter++
		}
	}

	compressed += com(prev, counter)
	if len(compressed) >= len(o) {
		return o
	}
	return compressed
}

func Test_compression(t *testing.T) {
	type args struct {
		orig string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{
				orig: "aaa",
			},
			want: "a3",
		},
		{
			name: "#2",
			args: args{
				orig: "aa",
			},
			want: "aa",
		},
		{
			name: "#3",
			args: args{
				orig: "abc",
			},
			want: "abc",
		},
		{
			name: "#4",
			args: args{
				orig: "abb",
			},
			want: "abb",
		},
		{
			name: "#5",
			args: args{
				orig: "aaaa",
			},
			want: "a4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression(tt.args.orig); got != tt.want {
				t.Errorf("compression() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_compression-8   	   62589	     17931 ns/op	    1327 B/op	      48 allocs/op
func Benchmark_compression(b *testing.B) {
	b.Run("#1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			compression("asdasdadaaasdasdadaaaa")
		}
	})

	b.Run("#2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			compression("abc")
		}
	})
}

func Fuzz_compression(f *testing.F) {
	f.Fuzz(func(t *testing.T, o string) {
		compression(o)
	})
}
