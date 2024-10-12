package sortAlgo

// PigeonholeSort sorts an array of int64 using the pigeonhole sorting algorithm
func PigeonholeSort(arr []int64) []int64 {
	minElement, maxElement := arr[0], arr[0]
	for _, num := range arr {
		if num < minElement {
			minElement = num
		}
		if num > maxElement {
			maxElement = num
		}
	}
	size := maxElement - minElement + 1

	// Create an array of pigeonholes (slice of slices), initialized to empty slices
	holes := make([][]int64, size)

	// Place each element in its respective hole
	for _, num := range arr {
		index := num - minElement
		holes[index] = append(holes[index], num)
	}

	// Collect the sorted elements
	idx := 0
	for i := int64(0); i < size; i++ {
		for _, num := range holes[i] {
			arr[idx] = num
			idx++
		}
	}

	return arr
}
