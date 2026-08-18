package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(37, problem37)
}

func isTruncatablePrime(p int) bool {
	place := 1
	n := 0
	for {
		n += (p % 10) * place
		if !util.IsPrime(n) || !util.IsPrime(p) {
			return false
		}
		place *= 10
		p /= 10
		if p == 0 {
			break
		}
	}
	return true
}

func problem37() {
	sum := 0
	count := 0
	for p := range util.Primes() {
		if p < 11 {
			continue
		}
		if isTruncatablePrime(p) {
			sum += p
			count++
			if count == 11 {
				break
			}
		}
	}
	fmt.Println(sum)
}
