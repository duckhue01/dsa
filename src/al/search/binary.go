package search

func Binary(arr []int, des int) bool {

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == des {
			return true
		}

		if des > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}
