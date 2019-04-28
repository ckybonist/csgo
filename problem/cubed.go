package problem

import (
	"fmt"
	"math"
)

type MyPair struct {
	x int
	y int
}
type Pairs = []*MyPair

func cube(n int) int {
	return int(math.Pow(float64(n), 3))
}

func (p *MyPair) equalTo(pp *MyPair) bool {
	return (p.x == pp.x && p.y == pp.y) || (p.x == pp.y && p.y == pp.x)
}
func (p *MyPair) reverse() *MyPair {
	return &MyPair{p.y, p.x}
}
func (p *MyPair) cubify() int {
	return cube(p.x) + cube(p.y)
}
func (p *MyPair) toString() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
func includesPair(pairs Pairs, pa *MyPair) bool {
	for _, pb := range pairs {
		if pa.equalTo(pb) {
			return true
		}
	}
	return false
}

// Find all possible (a, b, c, d) combination
// where a^3 + b^3 = c^3 + d^3.
// 0 < a, b, c, d < 1000
func Cubed() {
	// fmt.Println("# Find all possible (a, b, c, d) combination ")
	// fmt.Println("    where a^3 + b^3 = c^3 + d ^ 3  |  0 < a, b, c, d < 1000")
	LEN := 1000

	// Build a hashmap that have such format:
	// a^3 + b^3 : (a, b)
	buildMap := func() map[int]Pairs {
		result := make(map[int]Pairs)

		for i := 1; i <= LEN; i++ {
			for j := 1; j <= LEN; j++ {

				p := MyPair{i, j}

				key := cube(i) + cube(j)
				notExisted := !includesPair(result[key], &p)

				if notExisted {
					if result[key] != nil {
						result[key] = append(result[key], &p)
					} else {
						result[key] = Pairs{&p}
					}
				}
			}
		}

		// Delete the field only have one items
		for key, pairs := range result {
			if len(pairs) == 1 {
				delete(result, key)
			}
		}

		return result
	}

	dict := buildMap()
	failed := false

	// if pairs, ok := dict[1729]; ok {
	// 	fmt.Println("1729:")
	// 	for _, p := range pairs {
	// 		fmt.Println(p.toString())
	// 	}
	// }

	for key, pairs := range dict {
		special := ""
		if len(pairs) > 2 {
			special = "(special case)"
		}
		fmt.Printf("%d %v:   ", key, special)
		for _, p := range pairs {
			if p.cubify() != key {
				failed = true
			}
			fmt.Printf("%v, ", p.toString())
		}
		fmt.Println("")
	}

	if !failed {
		fmt.Println("Pass !!")
	}
	fmt.Println("end of process")
}
