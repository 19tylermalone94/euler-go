package utility

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactors(n int) []int {
	if IsPrime(n) {
		return []int{n}
	}
	primeFactors := []int{}
	for {
		if IsPrime(n) {
			primeFactors = append(primeFactors, n)
			break
		}
		for i := 2; i*i <= n; i++ {
			if IsPrime(i) && n%i == 0 {
				primeFactors = append(primeFactors, i)
				n /= i
				break
			}
		}
	}
	return primeFactors
}
