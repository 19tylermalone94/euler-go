package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(27, problem27)
}

func problem27() {
	maxPrimeCount := 0
	a, b := 0, 0
	for i := -999; i < 1000; i++ {
		for j := -1000; j <= 1000; j++ {
			n := 0
			for {
				if !util.IsPrime(n*n + i*n + j) {
					break
				}
				n++
			}
			if n > maxPrimeCount {
				maxPrimeCount = n
				a, b = i, j
			}
		}
	}
	fmt.Println(a * b)
}
