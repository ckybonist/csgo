package problem

import "fmt"

func FindIntersection() {
	fmt.Println("# Find intersection of two array:")
	a1 := []int{1, 3, 5, 7}
	a2 := []int{2, 1, 7}
	fmt.Println(a1)
	fmt.Println(a2)

	result := intersect(a1, a2)
	fmt.Println("Result: ", result)
}

func intersect(a []int, b []int) []int {
	var result []int
	dict := make(map[int]int)

	for _, e := range a {
		dict[e] = e
	}

	for _, bv := range b {
		if _, ok := dict[bv]; ok {
			result = append(result, bv)
		}
	}

	return result
}
