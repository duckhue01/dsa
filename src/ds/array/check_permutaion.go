package array

// given two string, write a method check if one is permutation of other

// with each element in a we loop through string b tc: O(n^2)
func checkPermutation1(stra, strb string) bool {

	if len(stra) == len(strb) {
		temp := make([]bool, len(stra))
		for i := 0; i < len(stra); i++ {
			ok := false
			for a := 0; a < len(strb); a++ {
				if stra[i] == strb[a] && !temp[a] {
					ok = true
					temp[a] = true
					break
				}
			}
			if !ok {
				return false
			}
		}
		return true
	}

	return false
}

// use hash table, assumption that we present string under ascii: O(n)
// if all character is presented in ascii
func checkPermutation3(a, b string) bool {
	table := make([]int, 127)

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			table[a[i]]++
		}

		for i := 0; i < len(b); i++ {
			table[b[i]]--
			if table[b[i]] < 0 {
				return false
			}
		}

		return true
	}

	return false
}

// sort string b then with each element in a we search that element in b: O(nlogn)
func checkPermutation2(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	str1 := []rune(a)
	str2 := []rune(b)
	stringQuickSort(str1, 0, len(a)-1)
	stringQuickSort(str2, 0, len(b)-1)

	for i := 0; i < len(a); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true

}

func stringQuickSort(str []rune, start int, end int) {

	if start < end {
		index := partition(str, start, end)
		stringQuickSort(str, start, index-1)
		stringQuickSort(str, index, end)
	}
}

func partition(str []rune, start int, end int) int {
	pivot := (start + (end-start)/2)

	for start <= end {

		for str[start] < str[pivot] {
			start++
		}

		for str[end] > str[pivot] {
			end--
		}

		if start <= end {
			temp := str[start]
			str[start] = str[end]
			str[end] = temp
			start++
			end--
		}
	}
	return start
}
