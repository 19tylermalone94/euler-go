package problems

import (
	"fmt"
	"math"
)

func init() {
	Register(45, problem45)
}

func IsHexagonal(n int) bool {
	sol := (1.0 + math.Sqrt(float64(1+8*n))) / 4.0
	return sol == float64(int(sol))
}

func problem45() {
	i := 286
	for {
		t := (i * (i + 1)) / 2
		if isPentagonalNumber(t) && IsHexagonal(t) {
			fmt.Println(t)
			return
		}
		i++
	}

}
