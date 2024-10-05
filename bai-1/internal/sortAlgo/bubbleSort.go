package sortAlgo

import "fmt"

func BubbleSort(arr []int64) []int64 {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("Processing i %d\n", i)
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println("Array arr: ", arr)
	return arr
}
