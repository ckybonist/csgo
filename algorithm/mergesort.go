package algorithm

import "fmt"

/*
func merge(arr []int, left int, mid int, right int) {
	L := mid - left + 1
	R := right - mid
	leftPart := make([]int, L)
	rightPart := make([]int, R)

	// Slice array by traditional way
	for i := 0; i < L; i++ {
		leftPart[i] = arr[left+i]
	}
	for i := 0; i < R; i++ {
		rightPart[i] = arr[mid+1+i]
	}

	l := 0
	r := 0
	k := left

	for l < L && r < R {
		if leftPart[l] <= rightPart[r] {
			arr[k] = leftPart[l]
			l++
		} else {
			arr[k] = rightPart[r]
			r++
		}
		k++
	}

	for l < L {
		arr[k] = leftPart[l]
		l++
		k++
	}
	for r < R {
		arr[k] = rightPart[r]
		r++
		k++
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", left, mid, right, leftPart, rightPart, arr)

}

func mergesort_old(arr []int, left int, right int) {
	if right > left {
		mid := left + (right-left)/2

		mergesort_old(arr, left, mid)
		mergesort_old(arr, mid+1, right)

		merge(arr, left, mid, right)
	}
}
*/

func divide(arr []int, mid int) ([]int, []int) {
	tmpL := arr[:mid]
	tmpR := arr[mid:]
	leftPart := make([]int, len(tmpL))
	rightPart := make([]int, len(tmpR))
	copy(leftPart, tmpL)
	copy(rightPart, tmpR)
	return leftPart, rightPart
}

func mergesort(arr []int) {
	N := len(arr)

	if N > 1 {
		mid := N / 2
		leftPart, rightPart := divide(arr, mid)
		L, R := len(leftPart), len(rightPart)

		mergesort(leftPart)
		mergesort(rightPart)

		l := 0
		r := 0
		k := 0

		for l < L && r < R {
			if leftPart[l] <= rightPart[r] {
				arr[k] = leftPart[l]
				l++
			} else {
				arr[k] = rightPart[r]
				r++
			}
			k++
		}

		for l < L {
			arr[k] = leftPart[l]
			l++
			k++
		}
		for r < R {
			arr[k] = rightPart[r]
			r++
			k++
		}
	}
}

func TestMergesort() {
	arr := []int{15, 3, 2, 7, 1, 4}

	fmt.Println("Merge Sort:")
	fmt.Printf("  Before: %v\n", arr)
	mergesort(arr)
	fmt.Printf("  After: %v", arr)
}
