package array

import (
	"testing"
)

type Pair struct {
	inp  string
	want bool
}

var str = []Pair{
	{"abc", true},
	{"ccccc", false},
	{"abcc", false},
	{"", true},
	{"aba", false},
	{"cccc", false},
	{"abbbc", false},
}

func TestIsUnique1(t *testing.T) {

	for i := 0; i < len(str); i++ {
		result := isUnique1(str[i].inp)
		if result != str[i].want {
			t.Fatalf("test case %d: func(%s) = %v but want match for %v", i, str[i].inp, result, str[i].want)
		}
	}

}

func TestIsUnique2(t *testing.T) {

	for i := 0; i < len(str); i++ {
		result := isUnique2(str[i].inp)
		if result != str[i].want {
			t.Fatalf("test case %d: func(%s) = %v but want match for %v", i, str[i].inp, result, str[i].want)
		}
	}

}


func TestIsUnique4(t *testing.T) {

	for i := 0; i < len(str); i++ {
		result := isUnique4(str[i].inp)
		if result != str[i].want {
			t.Fatalf("test case %d: func(%s) = %v but want match for %v", i, str[i].inp, result, str[i].want)
		}
	}

}
func TestIsUnique5(t *testing.T) {

	for i := 0; i < len(str); i++ {
		result := isUnique5(str[i].inp)
		if result != str[i].want {
			t.Fatalf("test case %d: func(%s) = %v but want match for %v", i, str[i].inp, result, str[i].want)
		}
	}

}


func TestIsUnique3(t *testing.T) {

	for i := 0; i < len(str); i++ {
		result := isUnique3(str[i].inp)
		if result != str[i].want {
			t.Fatalf("test case %d: func(%s) = %v but want match for %v", i, str[i].inp, result, str[i].want)
		}
	}

}