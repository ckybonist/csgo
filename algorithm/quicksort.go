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

func quicksort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
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
