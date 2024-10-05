package sortAlgo

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectionSort sorts an array using the Selection sort algorithm.
func SelectionSort(arr []int64) {
	fmt.Println("Initial array:", arr)

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
		fmt.Printf("Processed index %d (swapped with index %d): %v\n", i, minIndex, arr)
	}

	fmt.Println("Final sorted array:", arr)
}

func main() {
	const N = 10 // For easier tracking, reduce the size to 10 for demonstration
	numbers := make([]int64, N)
	rand.Seed(time.Now().UnixNano())

	// Generate random numbers
	for i := 0; i < N; i++ {
		numbers[i] = rand.Int63n(100) // Limit range to 0-99 for easier visualization
	}

	start := time.Now()
	SelectionSort(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Sorting completed in %v seconds.\n", elapsed.Seconds())
}
