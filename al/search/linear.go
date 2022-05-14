package search
// iterative implementation always better than recursive version because
// recursive version require extra memory space to store the context for previous 
// function so that in simple problem we should use iterative version
func Linear(arr []int, des int) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i] == des {
			return true
		}
	}

	return false
}
