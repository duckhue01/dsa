package sort

import (
	"testing"
)

type Pair struct {
	inp1 []int
	want []int
}

var pairs = []Pair{
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{3, 4, 5, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 4, 3, 2, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 3, 2, 4, 5}, []int{1, 2, 3, 4, 5}},
}

func TestInsetion(t *testing.T) {
	for i := 0; i < len(pairs); i++ {
		result := insertion(pairs[i].inp1)
		isOk := true
		if len(result) != len(pairs[i].want) {
			isOk = false;
		}
		for a := 0; a < len(result); a++ {
			if result[a] != pairs[i].want[a] {
				isOk = false
			}
		}

		if !isOk {
			t.Fatalf("test case %d: func(%v) = %v but want match for %v", i, pairs[i].inp1, result, pairs[i].want)
		}
	}
}
func TestMergeSort(t *testing.T) {
	for i := 0; i < len(pairs); i++ {
		result := mergeSort(pairs[i].inp1)
		isOk := true
		if len(result) != len(pairs[i].want) {
			isOk = false;
		}
		for a := 0; a < len(result); a++ {
			if result[a] != pairs[i].want[a] {
				isOk = false
			}
		}

		if !isOk {
			t.Fatalf("test case %d: func(%v) = %v but want match for %v", i, pairs[i].inp1, result, pairs[i].want)
		}
	}
}



func TestQuickSort(t *testing.T) {
	for i := 0; i < len(pairs); i++ {
		result := quickSort(pairs[i].inp1)
		isOk := true
		if len(result) != len(pairs[i].want) {
			isOk = false;
		}
		for a := 0; a < len(result); a++ {
			if result[a] != pairs[i].want[a] {
				isOk = false
			}
		}

		if !isOk {
			t.Fatalf("test case %d: func(%v) = %v but want match for %v", i, pairs[i].inp1, result, pairs[i].want)
		}
	}
}