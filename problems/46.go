package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(46, problem46)
}

func canDoIt(n int) bool {
	for p := range util.Primes() {
		if p >= n {
			break
		}
		for root := 1; ; root++ {
			y := p + 2*root*root
			if y > n {
				break
			}
			if y == n {
				return true
			}
		}
	}
	return false
}

func problem46() {
	for i := 9; ; i += 2 {
		if util.IsPrime(i) {
			continue
		}
		if !canDoIt(i) {
			fmt.Println(i)
			break
		}
	}
}
