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

// solution #1
func com(c byte, n int) string {
	if n == 1 {
		return string(c)
	}
	return fmt.Sprint(string(c), n)
}

func compression01(orig string) string {
	olen := len(orig)

	if olen == 0 {
		return orig
	}

	prev := orig[0]
	counter := 1
	cstring := ""
	clen := 0

	for i := 1; i < olen; i++ {
		if prev != orig[i] {
			res := com(prev, counter)
			cstring += res
			clen += len(res)
			if clen >= olen {
				return orig
			}
			prev = orig[i]
			counter = 1
		} else {
			counter++
		}
	}
	res := com(prev, counter)
	cstring += res
	clen += len(res)
	if clen >= olen {
		return orig
	}
	return cstring
}

func Test_compression01(t *testing.T) {
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
				orig: "aaabbcddd",
			},
			want: "a3b2cd3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression01(tt.args.orig); got != tt.want {
				t.Errorf("compression() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_compression01-8   	  186748	      5395 ns/op	     720 B/op	      38 allocs/op
func Benchmark_compression01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compression01("asdasdadaaasdasdadaaaa")
	}
}

// solution #2
type atom struct {
	key     byte
	counter int
}

func put(s *string, a *atom) int {
	*s = fmt.Sprint(*s, string(a.key), a.counter)
	return len(*s)
}

func compression02(o string) string {
	oLen := len(o)
	if oLen == 0 {
		return o
	}

	a := &atom{
		key:     o[0],
		counter: 1,
	}

	compressed := ""

	for i := 1; i < oLen; i++ {
		if a.key != o[i] {
			cLen := put(&compressed, a)
			if cLen >= oLen {
				return o
			}
			a = &atom{
				key:     o[i],
				counter: 1,
			}
		} else {
			a.counter++
		}
	}

	// put last element to c
	cLen := put(&compressed, a)
	if cLen >= oLen {
		return o
	}
	return compressed
}

func Test_compression02(t *testing.T) {
	type args struct {
		o string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{
				o: "aabbccccc",
			},
			want: "a2b2c5",
		},
		{
			name: "#2",
			args: args{
				o: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression02(tt.args.o); got != tt.want {
				t.Errorf("compression02() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_compression02-8   	   48879	     27357 ns/op	    1533 B/op	      60 allocs/op
func Benchmark_compression02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compression02("asdasdadaaasdasdadaaaa")
	}
}
