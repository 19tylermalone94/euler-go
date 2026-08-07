package problems

import "fmt"

func init() {
	Register(14, problem14)
}

func problem14() {
	dp := map[int]int{}
	maxCount := 0
	maxStart := 0
	for start := 1; start < 1000000; start++ {
		count := 0
		n := start
		for {
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
			count++
			val, ok := dp[n]
			if ok {
				count += val
				n = 1
			}
			if n == 1 {
				dp[start] = count
				if count > maxCount {
					maxCount = count
					maxStart = start
				}
				break
			}
		}
	}
	fmt.Println(maxStart, maxCount)
}
