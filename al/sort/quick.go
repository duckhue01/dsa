package sort

// [3, 4, 1, 2, 6, 2]
func quickSort(arr []int) []int {
	hepper(arr, 0, len(arr)-1)
	return arr
}

func hepper(arr []int, low int, high int) {

	if low < high {
		index := partition1(arr, low, high)
		hepper(arr, low, index-1)
		hepper(arr, index+1, high)
	}
}

func partition1(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for a := low; a < high; a++ {
		if arr[a] < pivot {
			i++
			temp := arr[a]
			arr[a] = arr[i]
			arr[i] = temp

		}
	}

	temp := arr[high]
	arr[high] = arr[i+1]
	arr[i+1] = temp

	return i + 1
}

