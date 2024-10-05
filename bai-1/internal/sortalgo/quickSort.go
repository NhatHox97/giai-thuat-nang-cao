package sortAlgo

import _ "fmt"

// QuickSort sorts an array of int64 using the Quick Sort algorithm
func QuickSort(arr []int64) {
	if len(arr) <= 1 {
		return
	}

	low, high := 0, len(arr)-1

	// Choose the pivot (in this case, we are using the middle element)
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Move the pivot to the rightmost position
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	// Partition the array
	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[low] = arr[low], arr[i]
			low++
		}
	}

	// Move the pivot to its final position
	arr[low], arr[high] = arr[high], arr[low]

	// Recursively apply QuickSort to the left and right partitions
	QuickSort(arr[:low])
	QuickSort(arr[low+1:])
}
