package bubble

// Sort implement bubble-sort algorithm.
func Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ { // *N
		for j := 0; j < len(arr)-1; j++ { // *(N-1)
			a, b := arr[j], arr[j+1]
			if a > b {
				arr[j], arr[j+1] = b, a
				a, b = arr[j], arr[j+1]
			} else {
				a, b = arr[j], arr[j+1]
			}
		}
	}
	return arr
}
