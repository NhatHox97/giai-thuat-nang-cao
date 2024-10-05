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
	// generateData()

	// Read files -> choose between Num1, Num2, Num3
	numbers, err := generator.ReadNumbers(Num3)
	if err != nil {
		fmt.Println("Can't find your find, please run generateData", err)
		return
	}

	// Calculate run time
	duration := calculateRunTime(func() {

		// Pass in your sorting algo here, all sorting algo inside package sortAlgo
		sortAlgo.ShellSort(numbers)
	})

	fmt.Printf("Sorting completed in %.2f seconds\n", duration)
}

/**
 * Calculate run time for the pass in function
 */
func calculateRunTime(runFunc func()) float64 {
	var start = time.Now().UnixMilli()
	runFunc() // Execute the passed-in function
	var end = time.Now().UnixMilli()
	var duration = float64(end-start) / 1000.0
	return duration
}

/**
 * Generate 3 files -> 1 mil, 10 mil, 100 mil
 */
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
