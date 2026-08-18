package problems

import (
	"fmt"
	"math"
)

func init() {
	Register(39, problem39)
}

func problem39() {
	maxP := 0
	maxSolutions := 0
	for p := 1; p <= 1000; p++ {
		numSolutions := 0
		for a := 1; a < p; a++ {
			for b := a + 1; b < p; b++ {
				cc := a*a + b*b
				c := int(math.Sqrt(float64(cc)))
				if c*c != cc {
					continue
				}
				if a+b+c == p {
					numSolutions++
				}
			}
		}
		if numSolutions > maxSolutions {
			maxP = p
			maxSolutions = numSolutions
		}
	}
	fmt.Println(maxP)
}
