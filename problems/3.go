package problems

import (
	"fmt"
	"slices"
)

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int) []int {
	res := []int{}
	for i := 2; i*i < n; i++ {
		if isPrime(i) && n%i == 0 {
			factor := n / i
			if isPrime(factor) {
				res = append(res, factor)
			} else {
				res = append(res, primeFactors(factor)...)
			}
		}
	}
	return res
}

func problem3() {
	fmt.Println(slices.Max(primeFactors(600851475143)))
}
