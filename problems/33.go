package problems

import (
	"fmt"
	"math"
)

func init() {
	Register(33, problem33)
}

func problem33() {
	numer := []int{}
	denom := []int{}
	for i := 1; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			a, b := i, j
			if a%10 != b/10 {
				continue
			}
			aa, bb := a/10, b%10
			x := float64(a) / float64(b)
			y := float64(aa) / float64(bb)
			if math.Abs(x-y) < 0.0001 {
				numer = append(numer, aa)
				denom = append(denom, bb)
			}
		}
	}
	prodNumer := 1
	prodDenom := 1
	for i := range numer {
		prodNumer *= numer[i]
		prodDenom *= denom[i]
	}
	fmt.Println(prodNumer, prodDenom)
}
