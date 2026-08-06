package problems

import (
	"euler-go/utility"
	"fmt"
	"slices"
)

func problem3() {
	fmt.Println(slices.Max(utility.PrimeFactors(600851475143)))
}
