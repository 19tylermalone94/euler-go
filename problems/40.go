package problems

import "fmt"

func init() {
	Register(40, problem40)
}

func problem40() {
	targets := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	k := 0
	product := 1
	i := 1
	n := 1
	for {
		if k >= len(targets) {
			break
		}
		nn := n
		digits := []int{}
		for nn > 0 {
			digits = append([]int{nn % 10}, digits...)
			nn /= 10
		}
		for _, digit := range digits {
			if i == targets[k] {
				product *= digit
				k++
				if k >= len(targets) {
					break
				}
			}
			i++
		}
		n++
	}
	fmt.Println(product)
}
