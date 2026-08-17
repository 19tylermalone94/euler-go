package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(34, problem34)
}

func problem34() {
	totalSum := 0
	for i := 3; i < 1000000; i++ {
		n := i
		sum := 0
		for n > 0 {
			sum += util.Factorial(n % 10)
			n /= 10
		}
		if sum == i {
			totalSum += sum
		}
	}
	fmt.Println(totalSum)
}
