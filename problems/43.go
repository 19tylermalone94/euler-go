package problems

import "fmt"

func init() {
	Register(43, problem43)
}

func hasDivProperty(n []int) bool {
	if len(n) != 10 {
		return false
	}
	targets := []int{2, 3, 5, 7, 11, 13, 17}
	k := 0
	for i := 1; i <= len(n)-3; i++ {
		num := 0
		for j := i; j < i+3; j++ {
			num = num*10 + n[j]
		}
		if num%targets[k] != 0 {
			return false
		}
		k++
	}
	return true
}

func sumSatPermutations(digits, perm []int) int {
	if len(perm) == 10 {
		if hasDivProperty(perm) {
			num := 0
			for _, d := range perm {
				num = num*10 + d
			}
			return num
		} else {
			return 0
		}
	}

	sum := 0
	for i := range digits {
		perm = append(perm, digits[i])
		digits = append(digits[0:i], digits[i+1:]...)
		sum += sumSatPermutations(digits, perm)
		digits = append(digits[0:i], append([]int{perm[len(perm)-1]}, digits[i:]...)...)
		perm = perm[0 : len(perm)-1]
	}
	return sum
}

func problem43() {
	fmt.Println(sumSatPermutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}))
}
