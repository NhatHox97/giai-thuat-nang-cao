package sortAlgo

import "fmt"

// InsertionSort sorts an array using the Insertion sort algorithm.
func InsertionSort(arr []int64) {
	fmt.Println("Initial array size: ", len(arr))

	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1], that are greater than temp,
		// to one position ahead of their current position
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp // Insert temp at the correct position

		// Print the state of the array after each insertion
		fmt.Printf("Processed element %d (value: %d): %v\n", i, temp, arr)
	}

	fmt.Println("Final sorted array:", arr)
}
