package problems

import (
	"euler-go/utility"
	"fmt"
	"slices"
)

func init() {
	Register(3, problem3)
}

func problem3() {
	fmt.Println(slices.Max(utility.PrimeFactors(600851475143)))
}
