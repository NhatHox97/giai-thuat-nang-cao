package sortAlgo

// ShellSort sorts an array using the Shell sort algorithm.
func ShellSort(arr []int64) {
	n := len(arr)
	gap := n / 2 // Start with a big gap, then reduce the gap
	// Do a gapped insertion sort for this gap size.
	for gap > 0 {
		// Pick all elements one by one and move them
		// to their correct position
		for i := gap; i < n; i++ {
			// Save the current element to be compared
			temp := arr[i]

			// Shift earlier gap-sorted elements up until the correct location
			// for arr[i] is found
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			// Put temp (the original arr[i]) in its correct location
			arr[j] = temp

			// Print the state of the array after each insertion
			//fmt.Printf("Processed element %d (value: %d) at index %d with gap %d: %v\n", i, temp, j, gap, arr)
		}
		gap /= 2 // Reduce the gap for the next iteration
		//fmt.Printf("Array after gap %d: %v\n", gap, arr)
	}
}
