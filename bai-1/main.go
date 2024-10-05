package main

import (
	"bai-1/internal/sortAlgo"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const size = 1000000
	testArr := make([]int64, size)

	// Populate the slice with random integers
	for i := 0; i < size; i++ {
		testArr[i] = rand.Int63n(1000000) // Random number between 0 and 999,999
	}

	duration := calculateRunTime(func() {
		sortAlgo.BubbleSort(testArr)
	})
	fmt.Printf("Durantion: %.3f\n", duration)
}

func calculateRunTime(runFunc func()) float64 {
	var start = time.Now().UnixMilli()
	runFunc()
	var end = time.Now().UnixMilli()
	var duration = float64(end-start) / 1000.0
	return duration
}
