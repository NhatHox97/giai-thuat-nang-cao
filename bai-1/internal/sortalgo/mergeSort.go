package sortAlgo

// MergeSort function sorts an array using the merge sort algorithm
func MergeSort(arr []int64) []int64 {
	// Base case: if the array has 1 or 0 elements, it is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle index
	mid := len(arr) / 2

	// Recursively split the array into two halves
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the two sorted halves
	return merge(left, right)
}

// Merge function merges two sorted arrays
func merge(left, right []int64) []int64 {
	result := []int64{}
	i, j := 0, 0

	// Compare elements from both arrays and merge them in sorted order
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements from both arrays (if any)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
