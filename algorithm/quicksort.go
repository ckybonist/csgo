package algorithm

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}

/**
 * Partition take last element as pivot, places
 * pivot at correct position, and places all smaller
 * element (smaller than pivot) to left of pivot,
 * all larger to right of pivot.
 *
 * Return: index of pivot
 * **/
func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	// Index of smaller element, start from -1
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = swap(arr[i], arr[j])
		}
	}

	arr[i+1], arr[high] = swap(arr[i+1], arr[high])

	return i + 1
}

// Tail recursion with O(Log n) extra space
func quicksort(arr []int, low int, high int) {
	for low < high {
		pivot := partition(arr, low, high)

		// Always choose the smaller part to do the recursion
		if pivot-low < high-pivot {
			quicksort(arr, low, pivot-1)
			low = pivot + 1 // iterate left part
		} else {
			quicksort(arr, pivot+1, high)
			high = pivot - 1
		}
	}
}

// TestQuicksort is a driver function
func TestQuicksort() {
	fmt.Println("Quicksort:")
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Printf("  Before: %v\n", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Printf("  After: %v", arr)
}
