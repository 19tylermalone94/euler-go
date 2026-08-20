package problems

import (
	"euler-go/util"
	"fmt"
	"maps"
)

func init() {
	Register(49, problem49)
}

func canSequence(start int, digits, perm []int, digitSet map[int]struct{}) bool {
	if len(perm) == 4 {
		num := 0
		for _, digit := range perm {
			num = num*10 + digit
		}
		if num <= start || !util.IsPrime(num) {
			return false
		}
		third := num + num - start
		thirdDigits := []int{}
		thirdDigitSet := make(map[int]struct{})
		t := third
		for t > 0 {
			thirdDigits = append(thirdDigits, t%10)
			thirdDigitSet[t%10] = struct{}{}
			t /= 10
		}
		fmt.Println(start, num, third)
		return util.IsPrime(third) && len(thirdDigits) == len(perm) && maps.Equal(thirdDigitSet, digitSet)
	}
	for i := range digits {
		perm = append(perm, digits[i])
		digits = append(digits[0:i], digits[i+1:]...)
		if canSequence(start, digits, perm, digitSet) {
			return true
		}
		digits = append(digits[0:i], append([]int{perm[len(perm)-1]}, digits[i:]...)...)
		perm = perm[0 : len(perm)-1]
	}
	return false
}

func problem49() {
	for p := range util.Primes() {
		if p < 1000 || p == 1487 {
			continue
		}
		digits := []int{}
		digitSet := make(map[int]struct{})
		pp := p
		for pp > 0 {
			d := pp % 10
			digits = append(digits, d)
			digitSet[d] = struct{}{}
			pp /= 10
		}
		if canSequence(p, digits, []int{}, digitSet) {
			fmt.Println(p)
			break
		}

		if p >= 10000 {
			break
		}
	}
}
