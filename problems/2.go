package problems

import "fmt"

func problem2() {
	sum := 0
	n2 := 0
	n1 := 1
	for {
		n := n2 + n1
		if n > 4000000 {
			break
		}
		if n%2 == 0 {
			sum += n
		}
		n2 = n1
		n1 = n
	}
	fmt.Println(sum)
}
