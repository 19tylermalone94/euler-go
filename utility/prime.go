package utility

func IsPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactors(n int) map[int]struct{} {
	res := make(map[int]struct{})
	for i := 2; i*i < n; i++ {
		if IsPrime(i) && n%i == 0 {
			factor := n / i
			if IsPrime(factor) {
				res[factor] = struct{}{}
			} else {
				for val := range PrimeFactors(factor) {
					res[val] = struct{}{}
				}
			}
		}
	}
	return res
}
