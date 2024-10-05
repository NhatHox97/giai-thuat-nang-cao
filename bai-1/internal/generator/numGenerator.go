package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GenerateNumbers creates a file with random numbers between 0 and 999999
func GenerateNumbers(filename string, count int) error {
	// Define the full path to save the file
	resourcesPath := filepath.Join("internal", "resources")

	// Check if the resources directory exists, and create it if it doesn't
	if _, err := os.Stat(resourcesPath); os.IsNotExist(err) {
		if err := os.MkdirAll(resourcesPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating resources directory: %w", err)
		}
	}

	fullPath := filepath.Join(resourcesPath, filename)

	// Create or open the file for writing
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < count; i++ {
		// Generate a random number between 0 and 999999
		num := rand.Int63n(1_000_000) // Use a proper random number generator
		if i > 0 {
			writer.WriteString(" ") // Space between numbers
		}
		_, err := writer.WriteString(strconv.FormatInt(num, 10))
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	writer.Flush() // Ensure all data is written to the file
	return nil
}

// ReadNumbers reads numbers from a file and returns them as a slice of int64
func ReadNumbers(filename string) ([]int64, error) {
	fullPath := filepath.Join("internal", "resources", filename)
	fmt.Println("Attempting to read from:", fullPath)

	// Read the entire file at once
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	parts := strings.Fields(string(data))
	numbers := make([]int64, len(parts))
	for i, part := range parts {
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error converting number to int64: %w", err)
		}
		numbers[i] = num
	}

	fmt.Println("Total numbers read from file:", len(numbers))
	return numbers, nil
}
