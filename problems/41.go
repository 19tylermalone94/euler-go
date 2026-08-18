package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(41, problem41)
}

func constructNumber(digits []int) int {
	num := 0
	for _, digit := range digits {
		num = num*10 + digit
	}
	return num
}

func maxPrimePermutation(n int, digits, perm []int) int {
	if len(perm) == n {
		return constructNumber(perm)
	}
	maxVal := 0
	for i := 0; i < len(digits); i++ {
		perm = append(perm, digits[i])
		digits = append(digits[0:i], digits[i+1:]...)
		val := maxPrimePermutation(n, digits, perm)
		if util.IsPrime(val) {
			maxVal = max(maxVal, val)
		}
		digits = append(digits[0:i], append([]int{perm[len(perm)-1]}, digits[i:]...)...)
		perm = perm[0 : len(perm)-1]
	}
	return maxVal
}

func problem41() {
	result := -1
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 9
	for n > 0 {
		maxPrimePerm := maxPrimePermutation(n, digits[0:n], []int{})
		if maxPrimePerm > 0 {
			result = maxPrimePerm
			break
		}
		n--
	}
	fmt.Println(result)
}
