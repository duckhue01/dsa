package array

// palindrome permutation: given a string, write a function to check if it
// is a permutation of palindrome. A palindrome is word or phase that is same
// forwards and backwards. A permutation is a rearrangement of letter.
// the palindrome does not need to be limited to just dictionary words.

// EX:
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco eta", etc.)

func palindromePermutation01(str []rune) bool {
	dic := make(map[rune]int)

	for _, v := range str {
		dic[v]++
	}

	if len(str)%2 == 0 {
		for _, v := range dic {
			if v%2 != 0 {
				return false
			}
		}
		return true
	} else {
		count := 0
		for _, v := range dic {
			if v%2 != 0 {
				count++
			}
			if count > 1 {
				return false
			}
		}
		return true
	}
}
