package sortAlgo

import (
	"sort"
)

// BucketSort function sorts an array using the bucket sort algorithm for int64
func BucketSort(arr []int64) []int64 {
	// Find the minimum and maximum values in the array
	minElement, maxElement := arr[0], arr[0]
	for _, num := range arr {
		if num > maxElement {
			maxElement = num
		} else if num < minElement {
			minElement = num
		}
	}

	// Create buckets
	n := len(arr)
	buckets := make([][]int64, n)

	// Distribute elements into the buckets
	for _, num := range arr {
		index := int((num - minElement) * int64(n-1) / (maxElement - minElement))
		buckets[index] = append(buckets[index], num)
	}

	// Sort each bucket and concatenate the result
	var sortedArr []int64
	for _, bucket := range buckets {
		sort.Slice(bucket, func(i, j int) bool { return bucket[i] < bucket[j] }) // Sort the individual bucket
		sortedArr = append(sortedArr, bucket...)
	}

	return sortedArr
}
