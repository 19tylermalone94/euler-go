package problems

import (
	"euler-go/utility"
	"fmt"
)

func init() {
	Register(23, problem23)
}

var abundantNums map[int]struct{} = make(map[int]struct{})

func isSumOfTwoAbundant(n int) bool {
	for num := range abundantNums {
		_, ok := abundantNums[n-num]
		if ok {
			return true
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
			abundantNums[i] = struct{}{}
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
