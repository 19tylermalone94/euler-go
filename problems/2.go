package problems

import "fmt"

func problem2() {
	sum := 0
	n2 := 0
	n1 := 1
	for {
		n := n2 + n1
		if n <= 4000000 {
			if n%2 == 0 {
				sum += n
			}
		} else {
			break
		}
		n2 = n1
		n1 = n
	}
	fmt.Println(sum)
}
