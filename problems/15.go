package problems

import "fmt"

func init() {
	Register(15, problem15)
}

func problem15() {
	n := 20
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	fmt.Println(dfs(0, 0, n, dp))
}

func dfs(i int, j int, n int, dp [][]int) int {
	if i == n && j == n {
		return 1
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	total := 0
	if i < n {
		total += dfs(i+1, j, n, dp)
	}
	if j < n {
		total += dfs(i, j+1, n, dp)
	}
	dp[i][j] = total
	return total
}
