package array

// #1: implement a function that determine each character in a string is unique

// use array as hash table to store number of time character appear
// time complexity: O(n) in worst case
// space complexity: O(1) in worst case
func isUnique1(str string) bool {

	// in case of ASCII (in case of unicode we need create bigger array)
	table := make([]bool, 256)

	for i := 0; i < len(str); i++ {
		if table[str[i]] {
			return false
		}
		table[str[i]] = true
	}

	return true
}

// we can reduce space by using int64 to present char (in case of charater in range a-Z)
// time complexity: O(n) in worst case
// space complexity: O(1) in worst case
func isUnique2(str string) bool {
	var val int64 = 0
	for i := 0; i < len(str); i++ {
		var temp int64 = 1 << (str[i] - 97)
		if (val & temp) != int64(0) {
			return false
		}
		val = val | temp
	}
	return true
}

// #2: what if we can not use additional data structure

// we can sort them check the neighbor if it same
// time complexity: O(nlogn) in worst case and the sort algorithm take O(nlogn) time complexity
// space complexity:
func isUnique3(str string) bool {

	return false
}

// we can use two loop to determind if string is unique
// time complexity: O(n^2)

func isUnique4(str string) bool {

	for i := 0; i < len(str)-1; i++ {
		for a := i + 1; a < len(str); a++ {
			if str[a] == str[i] {
				return false
			}
		}
	}

	return true
}

func isUnique5(str string) bool {
	return false
}
