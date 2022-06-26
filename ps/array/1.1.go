package array

// is unique: implement an algorithms to determine if a string a all unique character.
// what if you can't use additional data structures?

// idea:
// using map to check if a character occurred in map or not.
// we can specify whether we use ascii(128 characters) or unicode (256 characters)
func isUnique01(str []rune) bool {
	if len(str) > 128 {
		return false
	}
	charMap := make(map[rune]bool)
	for _, v := range str {
		if !charMap[v] {
			charMap[v] = true
		} else {
			return false
		}
	}
	return true
}

// wrapping up:
// time complexity: O(n)
// space complexity: O(1) because we have a fixed amount of character on alphabet
// the type of value of key in map use can use bool because it use just one bit to store value.

// idea:
// we can specify whether we use ascii(128 characters) or unicode (256 characters)
// we using array with pre defined length each for one character
func isUnique02(str []rune) bool {

	return true
}

// idea:
// we can reduce the space consuming by using bit manipulation because:
// + we can represent 64 character with int64(64bit)
// + the smallest addressable unit of computer is byte so when we use array or
// map it need at lest 8 bit for each value
func isUnique03(str []rune) bool {
	var a int64
	for _, v := range str {
		v = v - 96
		b := int64(1 << v)
		if a&b != 0 {
			return false
		}
		a = a ^ b
	}
	return true
}

// idea:
// we use nested loop to find the duplicate character.
func isUnique04(str []rune) bool {
	if len(str) > 128 {
		return false
	}
	search := func(e rune, idx int) bool {
		for i := idx + 1; i < len(str); i++ {
			if e == str[i] {
				return true
			}
		}
		return false

	}
	for i, v := range str {
		if search(v, i) {
			return false
		}
	}
	return true
}

// wrapping up:
// time complexity: O(n*n)
// space complexity: O(1)

// idea:
func isUnique05(str []rune) bool {

	return true
}
