package problem

import (
	"fmt"
	"strings"
)

func swap(x string, y string) (string, string) {
	return y, x
}


func heapPerm(s []string, size int, cap int, count *int) {
	if size == 1 {
		fmt.Println(s)
		(*count)++
	}

	for i := 0; i < size; i++ {
		heapPerm(s, size - 1, cap, count)

		
		if i < size - 1 {  // avoid unnecessary swaps when i == k-1
			if size % 2 == 1 {
				s[0], s[size - 1] = swap(s[0], s[size - 1])
			} else {
				s[i], s[size - 1] = swap(s[i], s[size - 1])
			}
		}
	}
}

func Perm() {
	s := strings.Split("abc", "")
	fmt.Printf("Permutations of %v:\n", s)
	count := 0
	heapPerm(s, len(s), len(s), &count)
	fmt.Println("Number of perms:", count)
}