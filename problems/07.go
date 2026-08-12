package problems

import (
	"euler-go/util"
	"fmt"
)

func init() {
	Register(7, problem7)
}

func problem7() {
	count := 0
	i := 2
	for {
		if util.IsPrime(i) {
			count++
			if count == 10001 {
				fmt.Println(i)
				return
			}
		}
		i++
	}
}
