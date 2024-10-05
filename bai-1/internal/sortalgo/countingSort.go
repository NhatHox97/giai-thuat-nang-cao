package sortAlgo

// CountingSort sorts an array of int64 using the counting sort algorithm
func CountingSort(arr []int64) []int64 {
	// Find the minElement and maxElement values
	minElement, maxElement := arr[0], arr[0]
	for _, num := range arr {
		if num < minElement {
			minElement = num
		}
		if num > maxElement {
			maxElement = num
		}
	}

	// Calculate the range
	rangeSize := maxElement - minElement + 1

	// Create a count array and output array
	count := make([]int64, rangeSize)
	output := make([]int64, len(arr))

	// Count the occurrences of each value
	for _, num := range arr {
		count[num-minElement]++
	}

	// Update count array to reflect positions
	for i := int64(1); i < rangeSize; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]-minElement]-1] = arr[i]
		count[arr[i]-minElement]--
	}

	return output
}
