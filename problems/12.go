package problems

import "fmt"

func init() {
	Register(12, problem12)
}

func problem12() {
	tnum := 0
	i := 1
	for {
		tnum += i
		if numDivisors(tnum) > 500 {
			break
		}
		i++
	}
	fmt.Println(tnum)
}

func numDivisors(n int) int {
	count := 0
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			count += 2
		}
	}
	return count
}
