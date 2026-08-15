package problems

import "fmt"

func init() {
	Register(28, problem28)
}

func problem28() {
	sum := 1
	i := 1
	s := 2
	for s <= 1000 {
		for range 4 {
			i += s
			sum += i
		}
		s += 2
	}
	fmt.Println(sum)
}
