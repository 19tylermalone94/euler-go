package problems

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Register(11, problem11)
}

func problem11() {
	n := 20
	k := 4
	data, err := os.ReadFile("input/11.txt")
	if err != nil {
		fmt.Println("Failed to read file")
		return
	}
	grid := [][]int{}
	for i, line := range strings.Split(string(data), "\n") {
		grid = append(grid, []int{})
		for str := range strings.SplitSeq(line, " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("failed to parse file")
				return
			}
			grid[i] = append(grid[i], num)
		}
	}

	di := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dj := []int{0, 1, 1, 1, 0, -1, -1, -1}

	maxPath := 0
	for i := k - 1; i < n-k-1; i++ {
		for j := k - 1; j < n-k-1; j++ {
			for d := range 8 {
				prod := 1
				for m := range 4 {
					prod *= grid[i+m*di[d]][j+m*dj[d]]
				}
				maxPath = max(maxPath, prod)
			}
		}
	}
	fmt.Println(maxPath)
}
