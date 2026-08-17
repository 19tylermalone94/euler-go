package problems

import (
	"fmt"
)

func init() {
	Register(32, problem32)
}

func problem32() {
	products := make(map[int]struct{})
	for i := 1; i < 100; i++ {
		for j := 100; j < 9999; j++ {
			digits := []int{}
			a, b, c := i, j, i*j
			for a > 0 {
				digits = append(digits, a%10)
				a /= 10
			}
			for b > 0 {
				digits = append(digits, b%10)
				b /= 10
			}
			for c > 0 {
				digits = append(digits, c%10)
				c /= 10
			}
			digitSet := make(map[int]struct{})
			for _, digit := range digits {
				digitSet[digit] = struct{}{}
			}
			_, ok := digitSet[0]
			if !ok && len(digits) == 9 && len(digitSet) == 9 {
				products[i*j] = struct{}{}
			}
		}
	}
	sum := 0
	for product := range products {
		sum += product
	}
	fmt.Println(sum)
}
