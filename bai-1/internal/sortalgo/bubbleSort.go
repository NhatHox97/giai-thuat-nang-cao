package sortAlgo

import "fmt"

func BubbleSort(arr []int64) []int64 {
	fmt.Println("Initial array size: ", len(arr))
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
		fmt.Printf("Processing i %d\n", i)
	}
	return arr
}