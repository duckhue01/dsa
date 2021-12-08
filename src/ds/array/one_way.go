package array

// One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.
// EXAMPLE
// -> true
// pales. pale -> true
// pale, pIe pale. bale -> true
// pale. bake -> false

func oneWay(a, b string) bool {

	i, k := 0, 0
	lena, lenb := len(a), len(b)
	count := 0

	if lena < lenb {
		for k < lenb && i < lena {
			if a[i] != b[k] {
				count++
				k++
			} else {
				k++
				i++
			}

		}
		count += lenb - k 
	}

	if lena > lenb {
		for k < lenb && i < lena {
			if a[i] != b[k] {
				count++
				i++
			} else {
				k++
				i++
			}
		}
		count += lena - i 
	}

	if lena == lenb {
		for k < lenb {
			if a[i] != b[k] {
				count++
			}
			k++
			i++
		}
	}

	return count < 2

}
