package problems

import (
	"euler-go/utility"
	"fmt"
)

func init() {
	Register(21, problem21)
}

func properDivisorSum(x int) int {
	properDivisors := utility.ProperDivisors(x)
	sum := 0
	for _, val := range properDivisors {
		sum += val
	}
	return sum
}

func isAmicable(x int) bool {
	y := properDivisorSum(x)
	return x != y && properDivisorSum(y) == x
}

func problem21() {
	n := 10000
	sum := 0
	for i := range n {
		if isAmicable(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
