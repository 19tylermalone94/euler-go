package util

import "math/big"

var BigOne *big.Int = big.NewInt(1)

func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func BigFactorial(n *big.Int) *big.Int {
	if n.Cmp(BigOne) <= 0 {
		return BigOne
	}
	result := new(big.Int)
	n1 := new(big.Int)
	n1 = n1.Sub(n, BigOne)
	return result.Mul(n, BigFactorial(n1))
}
