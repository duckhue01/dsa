package array

// import "fmt"

// input:  "duc khue 01", 11
// changed:"duc khue 01    "
// output: "duc%20khue%2001"

// time complexity(tx): O(n2)
// because every time we face the space we have to
// create new array from old array with tx is 0(n)
// so that in the worst case we have to copy n times
func URLify01(str string, size int) string {

	for i := 0; i < size; i++ {
		if str[i] == 32 {
			str = str[:i] + "%20" + str[i+1:]
		}
	}

	return str

}

// tx: O(n)
// because we just have to copy just one time
func URLify02(str []rune, size int) string {
	var sps int = 0

	for i := 0; i < size; i++ {
		if str[i] == 32 {
			sps++
		}
	}
	ix := size + sps*2
	str = append(str, make([]rune, ix-size)...)
	for i := size - 1; i >= 0; i-- {
		if str[i] == 32 {
			str[ix-1] = '0'
			str[ix-2] = '2'
			str[ix-3] = '%'
			ix -= 3
		} else {
			str[ix-1] = str[i]
			ix--
		}
	}

	return string(str)
}
