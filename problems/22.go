package problems

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func init() {
	Register(22, problem22)
}

func problem22() {
	data, err := os.ReadFile("input/22.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
	}
	names := strings.Split(string(data), ",")
	slices.Sort(names)

	sum := 0
	for i, name := range names {
		score := 0
		for _, char := range name {
			if char == '"' {
				continue
			}
			score += int(char) - 'A' + 1
		}
		sum += score * (i + 1)
	}
	fmt.Println(sum)
}
