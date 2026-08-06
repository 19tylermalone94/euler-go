package problems

import "fmt"

func problem1() {
	sum := 0
	for i := range 1000 {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
