package sortAlgo

// InterchangeSort function sorts an array using the interchange sort algorithm
func InterchangeSort(arr []int64) []int64 {
	n := len(arr)

	// Loop over all elements in the array
	for i := 0; i < n-1; i++ {
		// Compare with all elements after i
		for j := i + 1; j < n; j++ {
			// Swap if the current element is greater than the compared element
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
