package array

// Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palin-
// drome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
// is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
// EXAMPLE
// Input: Tact Coa
// Output: True (permutations: "taco cat". "atco cta". etc.)
// Hints: # 106, #121, #134, #136

// time complexity : O(n)
func palindrome(str string) bool {

	table := make(map[rune]int)
	dup := 0
	count := 0

	for _, v := range str {
		if v != 32 && v > -1 {
			table[v]++

			if table[v]%2 == 0 {
				dup--
			} else {
				dup++
			}
			count++
		}

	}


	return dup  <= 1
}

// optimized palindrome
func palindrome01(str string) bool {

	var bt = 0

	for _, v := range str {
		v = (v - 'a')
		if v > -1 {
			mask := 1 << v
			if mask&bt == 0 {
				bt |= mask
			} else {
				bt &= ^mask
			}
		}
	}

	return bt == 0 || (bt & (bt - 1)) == 0
}
