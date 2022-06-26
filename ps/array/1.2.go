package array

// check permutation: given two strings, write a method to decide if one is
// a permutation of the other.

// idea:
// gradually add each element of add each element of array 1 to map at the same time add it by one
// do the same thing with array 2
// compare length of two map and each element in that map.
func checkPermutation01(str1 []rune, str2 []rune) bool {
	if len(str1) != len(str2) {
		return false
	}
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, v := range str1 {
		map1[v]++
	}
	for _, v := range str2 {
		map2[v]++
	}
	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}

	return true
}

// wrapping up:
// time complexity: O(n)
// space complexity: O(n)

func checkPermutation02(str1 []rune, str2 []rune) bool {
	if len(str1) != len(str2) {
		return false
	}
	temp := [128]int{}

	for _, v := range str1 {
		temp[v]++
	}
	for _, v := range str2 {
		if i := temp[v]; i < 0 {
			return false
		}
	}
	return true
}

// wrapping up:
// time complexity: O(n)
// space complexity: O(n)
