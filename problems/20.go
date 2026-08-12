package problems

import (
	"fmt"
	"math/big"
)

var one *big.Int = big.NewInt(1)

func init() {
	Register(20, problem20)
}

func factorial(n *big.Int) *big.Int {
	if n.Cmp(one) <= 0 {
		return one
	}
	result := new(big.Int)
	n1 := new(big.Int)
	n1 = n1.Sub(n, one)
	return result.Mul(n, factorial(n1))
}

func problem20() {
	n := big.NewInt(100)
	fn := factorial(n).String()
	sum := 0
	for i := range fn {
		sum += int(fn[i]) - '0'
	}
	fmt.Println(sum)
}
