package search

import (
	"testing"
)

type Pair struct {
	inp1 []int
	inp2 int
	want bool
}

var pairs = []Pair{
	{[]int{1, 2, 3, 4, 5}, 2, true},
	{[]int{1, 2, 3, 4, 5}, 12, false},
	{[]int{1, 2, 3, 4, 5}, 5, true},
	{[]int{1, 2, 3, 4, 5}, 1, true},
	{[]int{1, 2, 3, 4, 5}, 0, false},
}

func TestLinear(t *testing.T) {

	for i := 0; i < len(pairs); i++ {
		result := Linear(pairs[i].inp1, pairs[i].inp2)
		if result != pairs[i].want {
			t.Fatalf("test case %d: func(%v, %d) = %v but want match for %v", i, pairs[i].inp1, pairs[i].inp2, result, pairs[i].want)
		}
	}
}

func TestBinary(t *testing.T) {
	for i := 0; i < len(pairs); i++ {
		result := Binary(pairs[i].inp1, pairs[i].inp2)
		if result != pairs[i].want {
			t.Fatalf("test case %d: func(%v, %d) = %v but want match for %v", i, pairs[i].inp1, pairs[i].inp2, result, pairs[i].want)
		}
	}
}
