package problems

import (
	"euler-go/utility"
	"fmt"
)

func problem7() {
	count := 0
	i := 2
	for {
		if utility.IsPrime(i) {
			count++
			if count == 10001 {
				fmt.Println(i)
				return
			}
		}
		i++
	}
}
