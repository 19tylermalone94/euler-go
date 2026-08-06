package problems

import "fmt"

func init() {
	Register(1, problem1)
}

func problem1() {
	sum := 0
	for i := range 1000 {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
