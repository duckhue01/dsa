package array

import (
	"testing"
)

// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
// Hints:#51, #100

// idea: rotate matrix in circular meaner

func RotateMatrixUsingTemp(m [][][4]byte) [][][4]byte {
	res := make([][][4]byte, len(m))
	for i := range res {
		res[i] = make([][4]byte, len(m))
	}

	for i, b := 0, len(m)-1; i < len(m) && b >= 0; i, b = i+1, b-1 {
		for j, a := 0, 0; j < len(m) && a < len(m); j, a = j+1, a+1 {
			res[a][b] = m[i][j]
		}
	}

	return res
}

func RotateMatrixWithoutTemp(m [][]int) [][]int {
	return [][]int{}
}

func TestXxx(t *testing.T) {
	t.Log(RotateMatrixUsingTemp([][][4]byte{
		{
			[4]byte{1, 2, 3, 4}, [4]byte{1, 2, 3, 4}, [4]byte{1, 2, 3, 4}},
		{[4]byte{1, 2, 3, 4}, [4]byte{1, 2, 3, 4}, [4]byte{1, 2, 3, 4}},
		{[4]byte{1, 2, 3, 4}, [4]byte{1, 2, 3, 4}, [4]byte{1, 2, 3, 4}}}))
}
