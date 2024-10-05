package sortAlgo

// RadixSort sorts the array of int64 using the radix sort algorithm
func RadixSort(arr []int64) []int64 {
	// Get the maximum number to determine the number of digits
	maxElement := arr[0]
	for _, num := range arr {
		if num > maxElement {
			maxElement = num
		}
	}

	// Perform sorting for each digit
	for exp := int64(1); maxElement/exp > 0; exp *= 10 {
		n := len(arr)
		output := make([]int64, n) // Output array
		count := make([]int64, 10) // Count array for digits 0-9

		// Count occurrences of each digit
		for i := 0; i < n; i++ {
			index := (arr[i] / exp) % 10
			count[index]++
		}

		// Update count[i] to contain the actual position of this digit in output[]
		for i := int64(1); i < 10; i++ {
			count[i] += count[i-1]
		}

		// Build the output array
		for i := n - 1; i >= 0; i-- {
			index := (arr[i] / exp) % 10
			output[count[index]-1] = arr[i]
			count[index]--
		}

		// Copy the output array back to arr, so that arr contains sorted numbers
		for i := 0; i < n; i++ {
			arr[i] = output[i]
		}
	}

	return arr
}
