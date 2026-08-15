package util

import (
	"iter"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Primes() iter.Seq[int] {
	return func(yield func(int) bool) {
		n := 2
		for {
			if IsPrime(n) && !yield(n) {
				return
			}
			n++
		}
	}
}

func PrimeFactors(n int) []int {
	if !IsPrime(n) {
		for k := range Primes() {
			if n%k == 0 {
				return append([]int{k}, PrimeFactors(n/k)...)
			}
		}
	}
	return []int{n}
}
