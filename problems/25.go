package problems

import (
	"euler-go/util"
	"fmt"
	"math/big"
)

func init() {
	Register(25, problem25)
}

func problem25() {
	n1 := big.NewInt(1)
	n2 := big.NewInt(1)
	i := 3
	for {
		n := new(big.Int).Add(n1, n2)
		if util.BigDigitCount(n) == 1000 {
			fmt.Println(i)
			break
		}
		n2 = n1
		n1 = n
		i++
	}
}
