package main

import (
	"bai-1/internal/generator"
	"bai-1/internal/sortalgo"
	"fmt"
	"time"
)

var Num1 = "num1.txt"
var Num2 = "num2.txt"
var Num3 = "num3.txt"

func main() {

	// Run this if you want to re-generate file
	//generateData()

	// Read files -> choose between Num1, Num2, Num3
	numbers, err := generator.ReadNumbers(Num1)
	if err != nil {
		fmt.Println("Can't find your files, please run generateData", err)
		return
	}

	// Calculate run time
	duration := calculateRunTime(func() {
		// Pass in your sorting algo here, all sorting algo inside package sortAlgo
		// Example: calculate runtime for quick sort
		sortAlgo.QuickSort(numbers)
	})

	fmt.Printf("Sorting completed in %.3f seconds\n", duration)
}

// calculateRunTime: calculate the runtime of the pass-in function
func calculateRunTime(runFunc func()) float64 {
	var start = time.Now().UnixMilli()
	runFunc() // Execute the passed-in function
	var end = time.Now().UnixMilli()
	var duration = float64(end-start) / 1000.0
	return duration
}

// generateData: generate 3 files for testing
func generateData() {
	err1 := generator.GenerateNumbers(Num1, 1_000_000)
	if err1 != nil {
		return
	}
	err2 := generator.GenerateNumbers(Num2, 10_000_000)
	if err2 != nil {
		return
	}
	err3 := generator.GenerateNumbers(Num3, 100_000_000)
	if err3 != nil {
		return
	}
}
