package problems

import (
	"fmt"
	"math/big"
)

func init() {
	Register(29, problem29)
}

func problem29() {
	nums := make(map[string]struct{})
	for i := 2; i <= 100; i++ {
		for j := 2; j <= 100; j++ {
			a := big.NewInt(int64(i))
			b := big.NewInt(int64(j))
			nums[new(big.Int).Exp(a, b, nil).String()] = struct{}{}
		}
	}
	fmt.Println(len(nums))
}
