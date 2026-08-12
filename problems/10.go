package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(10, problem10)
}

func problem10() {
	sum := 0
	i := 2
	for {
		if i >= 2000000 {
			break
		}
		if util.IsPrime(i) {
			sum += i
		}
		i++
	}
	fmt.Println(sum)
}
