package sort

func divide(arr []int, left int, right int) {
	if left == right {
		return
	} else {
		mid := left + (right-left)/2
		divide(arr, left, mid)
		divide(arr, mid+1, right)
		merge(arr, left, mid, right)

	}
}

func merge(arr []int, left int, mid int, right int) {
	size1, size2 := mid-left+1, right-mid

	arr1, arr2 := make([]int, size1), make([]int, size2)

	for i := 0; i < size1; i++ {
		arr1[i] = arr[left+i]
	}
	for i := 0; i < size2; i++ {
		arr2[i] = arr[mid+1+i]
	}

	a, b, k := 0, 0, left

	for (a < size1) && (b < size2) {
		if arr1[a] <= arr2[b] {
			arr[k] = arr1[a]
			a++
		} else {
			arr[k] = arr2[b]
			b++
		}
		k++
	}

	for a < size1 {
		arr[k] = arr1[a]
		a++
		k++
	}

	for b < size2 {
		arr[k] = arr2[b]
		b++
		k++
	}

}

func mergeSort(arr []int) []int {

	divide(arr, 0, len(arr)-1)

	return arr
}
