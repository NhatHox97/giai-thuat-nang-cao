package sortAlgo

import (
	"fmt"
)

// SelectionSort sorts an array using the Selection sort algorithm.
func SelectionSort(arr []int64) {
	fmt.Println("Initial array:", len(arr))

	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Assume the minimum is the first element
		minIndex := i
		for j := i + 1; j < n; j++ {
			// Find the minimum element in the unsorted portion
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]

		// Print the state of the array after each swap
	}
}
