package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(35, problem35)
}

func isCircularPrime(n int) bool {
	numDigits := util.DigitCount(n)
	curr := n
	for {
		if !util.IsPrime(curr) {
			return false
		}
		lastDigit := curr % 10
		rest := curr / 10
		rotated := lastDigit*util.Pow(10, numDigits-1) + rest
		if rotated == n {
			break
		}
		curr = rotated
	}
	return true
}

func problem35() {
	count := 0
	for i := range 1000000 {
		if isCircularPrime(i) {
			count++
		}
	}
	fmt.Println(count)
}
