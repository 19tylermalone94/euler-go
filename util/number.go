package util

import "math/big"

func DigitCount(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func BigDigitCount(n *big.Int) int {
	return len(n.String())
}
