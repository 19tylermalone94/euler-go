package util

func ProperDivisors(x int) []int {
	properDivisors := []int{}
	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			properDivisors = append(properDivisors, i)
		}
	}
	return properDivisors
}
