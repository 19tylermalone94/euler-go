package problems

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Register(13, problem13)
}

func problem13() {
	n := 100
	m := 50
	data, err := os.ReadFile("input/13.txt")
	if err != nil {
		fmt.Println("Failed to read file")
		return
	}
	grid := [][]int{}
	for i, line := range strings.Split(string(data), "\n") {
		grid = append(grid, []int{})
		for str := range strings.SplitSeq(line, "") {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("failed to parse int", str)
				return
			}
			grid[i] = append(grid[i], num)
		}
	}

	solution := []int{}
	carry := 0
	for j := m - 1; j >= 0; j-- {
		sum := 0
		for i := range n {
			sum += grid[i][j]
		}
		sum += carry
		carry = sum / 10
		fmt.Println(sum, carry)
		solution = append([]int{sum % 10}, solution...)
	}
	if carry > 0 {
		solution = append([]int{carry}, solution...)
	}
	b := strings.Builder{}
	for _, n := range solution {
		b.WriteString(strconv.Itoa(n))
	}
	fmt.Println(b.String())
}
