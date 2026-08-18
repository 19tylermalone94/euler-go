package problems

import "fmt"

func init() {
	Register(38, problem38)
}

func problem38() {
	maxPandigital := 0
	for i := 1; i < 10000; i++ {
		digitSet := make(map[int]struct{})
		digits := []int{}
		n := 1
		for {
			if len(digits) >= 9 {
				break
			}
			d := []int{}
			k := i * n
			for k > 0 {
				d = append([]int{k % 10}, d...)
				digitSet[k%10] = struct{}{}
				k /= 10
			}
			digits = append(digits, d...)
			n++
		}
		_, ok := digitSet[0]
		if len(digits) == 9 && len(digitSet) == 9 && !ok {
			pandigital := 0
			for _, digit := range digits {
				pandigital = pandigital*10 + digit
			}
			maxPandigital = max(maxPandigital, pandigital)
		}
	}
	fmt.Println(maxPandigital)
}
