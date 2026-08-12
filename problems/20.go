package problems

import (
	"euler-go/util"
	"fmt"
	"math/big"
)

func init() {
	Register(20, problem20)
}

func problem20() {
	n := big.NewInt(100)
	fn := util.Factorial(n).String()
	sum := 0
	for i := range fn {
		sum += int(fn[i]) - '0'
	}
	fmt.Println(sum)
}
