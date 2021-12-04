package array

import "testing"


type Triad struct{
	a string
	b string
	c bool
}

var cases = []Triad{
	{"abcccc", "cccbca", true},
	{"abc", "bca", true},
	{"abc", "bcac", false},
	{"abca", "bcaa", true},
	{"abcab", "bcaba", true},
	{"abcaa", "bca", false},
	{"abc", "bcadd", false},
	{"aaa", "bca", false},

}



func TestCheckPermutation1(t *testing.T) {

	for i := 0; i < len(cases); i++ {
		re := checkPermutation1(cases[i].a, cases[i].b)
		if re != cases[i].c {
			t.Fatalf("test case %d: func(%s, %s) = %v but want match for %v",i, cases[i].a, cases[i].b, re, cases[i].c )
		}
	}
}

func TestCheckPermutation2(t *testing.T) {

	for i := 0; i < len(cases); i++ {
		re := checkPermutation2(cases[i].a, cases[i].b)
		if re != cases[i].c {
			t.Fatalf("test case %d: func(%s, %s) = %v but want match for %v",i, cases[i].a, cases[i].b, re, cases[i].c )
		}
	}
}


func TestCheckPermutation3(t *testing.T) {

	for i := 0; i < len(cases); i++ {
		re := checkPermutation3(cases[i].a, cases[i].b)
		if re != cases[i].c {
			t.Fatalf("test case %d: func(%s, %s) = %v but want match for %v",i, cases[i].a, cases[i].b, re, cases[i].c )
		}
	}
}
