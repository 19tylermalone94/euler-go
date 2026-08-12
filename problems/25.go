package problems

import (
	"fmt"
	"math/big"
)

func init() {
	Register(25, problem25)
}

func digitCount(n *big.Int) int {
	return len(n.String())
}

func problem25() {
	n1 := big.NewInt(1)
	n2 := big.NewInt(1)
	i := 3
	for {
		n := new(big.Int).Add(n1, n2)
		if digitCount(n) == 1000 {
			fmt.Println(i)
			break
		}
		n2 = n1
		n1 = n
		i++
	}
}
