package problems

import (
	"fmt"
	"math/big"
)

func init() {
	Register(48, problem48)
}

func problem48() {
	sum := new(big.Int)
	for i := 1; i <= 1000; i++ {
		sum = new(big.Int).Add(sum, new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil))
	}
	digits := sum.String()
	fmt.Println(digits[len(digits)-10:])
}
