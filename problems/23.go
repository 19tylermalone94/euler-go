package problems

import (
	"euler-go/utility"
	"fmt"
)

func init() {
	Register(23, problem23)
}

var abundantNums []int = []int{}

func isSumOfTwoAbundant(n int) bool {
	for _, val1 := range abundantNums {
		if val1 >= n {
			break
		}
		for _, val2 := range abundantNums {
			if val1+val2 > n {
				break
			}
			if val1+val2 == n {
				return true
			}
		}
	}
	return false
}

func problem23() {
	for i := range 28124 {
		properDivisors := utility.ProperDivisors(i)
		sum := 0
		for _, val := range properDivisors {
			sum += val
		}
		if sum > i {
			abundantNums = append(abundantNums, i)
		}
	}
	sum := 0
	for i := range 28124 {
		if !isSumOfTwoAbundant(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
