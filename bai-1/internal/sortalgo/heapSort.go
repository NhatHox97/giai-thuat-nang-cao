package sortAlgo

// HeapSort sorts an array using heap sort algorithm
func HeapSort(arr []int64) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		processHeapLogic(arr, int64(n), int64(i))
	}

	// Extract elements from the heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// Call max processHeapLogic on the reduced heap
		processHeapLogic(arr, int64(i), 0)
	}
}

// Subtree rooted with node i, which is an index in arr[]
// n is the size of the heap
func processHeapLogic(arr []int64, n int64, i int64) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // left = 2*i + 1
	right := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than the largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap

		// Recursively processHeapLogic the affected sub-tree
		processHeapLogic(arr, n, largest)
	}
}
