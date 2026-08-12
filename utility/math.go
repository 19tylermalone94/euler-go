package utility

import "math/big"

var BigOne *big.Int = big.NewInt(1)

func Factorial(n *big.Int) *big.Int {
	if n.Cmp(BigOne) <= 0 {
		return BigOne
	}
	result := new(big.Int)
	n1 := new(big.Int)
	n1 = n1.Sub(n, BigOne)
	return result.Mul(n, Factorial(n1))
}
