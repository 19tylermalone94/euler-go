package problems

import (
	"fmt"
)

func init() {
	Register(47, problem47)
}

func problem47() {
	const limit = 200000
	nf := make([]int, limit)
	for p := 2; p < limit; p++ {
		if nf[p] == 0 {
			for m := p; m < limit; m += p {
				nf[m]++
			}
		}
	}
	run := 0
	for n := 2; n < limit; n++ {
		if nf[n] == 4 {
			if run++; run == 4 {
				fmt.Println(n - 3)
				return
			}
		} else {
			run = 0
		}
	}
}
