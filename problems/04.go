package problems

import "fmt"

func init() {
	Register(4, problem4)
}

func isPalindromeNumber(n int) bool {
	temp := n
	reverse := 0
	for {
		if temp <= 0 {
			break
		}
		end := temp % 10
		reverse = reverse*10 + end
		temp /= 10
	}
	return reverse == n
}

func problem4() {
	max := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 0; j-- {
			product := i * j
			if product > max && isPalindromeNumber(product) {
				max = product
			}

			if product < max {
				break
			}
		}
	}
	fmt.Println(max)
}
