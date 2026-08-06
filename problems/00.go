package problems

import "fmt"

func init() {
	Register(0, problem0)
}

func problem0() {
	sum := 0
	for i := range 587000 {
		square := i * i
		if square%2 == 1 {
			sum += square
		}
	}
	fmt.Println(sum)
}
