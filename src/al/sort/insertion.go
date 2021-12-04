package sort

func insertion(arr []int) []int {

	for i := 1; i < len(arr); i++ {

		a := i - 1
		key := arr[i]
		for a >= 0 && arr[a] > key {

			arr[a+1] = arr[a]
			a--
		}
		arr[a+1] = key

	}
	return arr
}
