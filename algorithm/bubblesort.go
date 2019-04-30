package algorithm

import "fmt"

func bubbleSort(arr []int) {
	isSorted := false
	lastUnsorted := len(arr) - 1

	for !isSorted {
		isSorted = true
		for i := 0; i < lastUnsorted; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = swap(arr[i], arr[i+1])
				isSorted = false
			}
		}

		lastUnsorted--
	}
}

// Bubblesort is a driver function
func Bubblesort() {
	arr := []int{15, 3, 2, 7, 1, 4}

	fmt.Println("Bubble Sort:")
	fmt.Printf("  Before: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("  After: %v", arr)
}
