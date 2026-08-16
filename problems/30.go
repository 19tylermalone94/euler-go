package problems

import (
	"fmt"
	"math"
	"math/big"
)

func init() {
	Register(30, problem30)
}

func problem30() {
	totalSum := new(big.Int)
	for i := 2; i < 1000000; i++ {
		sum := 0
		n := i
		for n > 0 {
			sum += int(math.Pow(float64(n%10), 5))
			n /= 10
		}
		if sum == i {
			totalSum = new(big.Int).Add(totalSum, big.NewInt(int64(i)))
		}
	}
	fmt.Println(totalSum)
}
