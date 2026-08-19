package problems

import (
	"fmt"
	"math"
)

func init() {
	Register(44, problem44)
}

func isPentagonalNumber(n int) bool {
	sol := (0.5 + math.Sqrt(0.25+float64(6*n))) / 3.0
	if sol > 0 && sol == float64(int(sol)) {
		return true
	}
	return false
}

func problem44() {
	pNums := []int{}
	p := 1
	d := 4
	for {
		for _, other := range pNums {
			sum := p + other
			dif := p - other
			if isPentagonalNumber(sum) && isPentagonalNumber(dif) {
				fmt.Println(p - other)
				return
			}
		}
		pNums = append(pNums, p)
		p += d
		d += 3
	}
}
