package problems

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Register(18, problem18)
}

func problem18() {
	data, err := os.ReadFile("input/18.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}
	grid := [][]int{}
	for i, line := range strings.Split(string(data), "\n") {
		grid = append(grid, []int{})
		for _, char := range strings.Split(line, " ") {
			num, err := strconv.Atoi(char)
			if err != nil {
				fmt.Println("Failed to parse int", char)
				return
			}
			grid[i] = append(grid[i], num)
		}
	}
	n := len(grid)
	for i := n - 2; i >= 0; i-- {
		for j := range grid[i] {
			down := grid[i+1][j]
			downRight := grid[i+1][j+1]
			grid[i][j] += max(down, downRight)
		}
	}
	fmt.Println(grid[0][0])
}
