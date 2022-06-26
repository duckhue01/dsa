package array

// URLify: write a method to replace all spaces in a string with '%20'. You may
// assume that the string has sufficient space at the end to hold the additional
// characters, and that you ware given the "true" length of the string
// example:
// Input: "Mr John Smith ", 13
// Output: "Mr%20John%20Smith"
func URLify01(str []rune, length int) []rune {
	sps := 0

	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			sps++
		}
	}

	newStr := make([]rune, length+(2*sps))
	x := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			newStr[x] = '%'
			newStr[x+1] = '2'
			newStr[x+2] = '0'
			x += 3
		} else {
			newStr[x] = str[i]
			x++
		}
	}
	return newStr
}

// Wrapping up:
// time complexity: O(n)
// space complexity: O(n)
