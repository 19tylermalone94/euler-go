package problems

import "fmt"

func init() {
	Register(31, problem31)
}

var coins []int = []int{1, 2, 5, 10, 20, 50, 100, 200}

func problem31() {
	n := 200
	fmt.Println(combos(0, n, 0))
}

func combos(i, n, sum int) int {
	if sum == n {
		return 1
	}
	if sum > n || i >= len(coins) {
		return 0
	}
	return combos(i, n, sum+coins[i]) + combos(i+1, n, sum)
}
