package problems

import (
	"euler-go/utility"
	"fmt"
	"maps"
	"slices"
)

func problem3() {
	fmt.Println(slices.Max(slices.Collect(maps.Keys(utility.PrimeFactors(600851475143)))))
}
