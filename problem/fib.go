package problem

import "fmt"

func memoFib(n int, memory map[int]int) int {
	if v, ok := memory[n]; ok {
		return v
	}
	memory[n] = memoFib(n-1, memory) + memoFib(n-2, memory)

	return memory[n]
}

// Fib number by tail recursion
func tailFib(n int, prev int, curr int) int {
	if n == 0 {
		return prev
	}
	if n == 1 {
		return curr
	}

	fmt.Printf("%d ", curr)

	return tailFib(n-1, curr, prev+curr)
}

func Fib(n int) int {
	num := tailFib(n, 0, 1)
	fmt.Printf("[%d]\n", num)

	// memory := map[int]int{
	// 	0: 0,
	// 	1: 1,
	// 	2: 1,
	// }
	// num := memoFib(n, memory)

	return num
}
