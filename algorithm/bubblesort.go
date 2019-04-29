package algorithm

import "fmt"

func bubbleSort(arr []int) {
	isSorted := false
	lastUnsorted := len(arr) - 1

	for !isSorted {
		isSorted = true
		for i := 0; i < lastUnsorted; i++ {
			if arr[i] > arr[i+1] {
				x, y := swap(arr[i], arr[i+1])
				arr[i] = x
				arr[i+1] = y
				isSorted = false
			}
		}

		lastUnsorted--
	}
}

// TestBubblesort is a driver function for bubble sort
func TestBubblesort() {
	arr := []int{15, 3, 2, 7, 1, 4}

	fmt.Println("Bubble Sort:")
	fmt.Printf("  Before: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("  After: %v", arr)
}
