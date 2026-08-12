package problems

import (
	"euler-go/util"
	"fmt"
	"slices"
)

func init() {
	Register(3, problem3)
}

func problem3() {
	fmt.Println(slices.Max(util.PrimeFactors(600851475143)))
}
