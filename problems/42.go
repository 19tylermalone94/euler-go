package problems

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func init() {
	Register(42, problem42)
}

func isTriangleNumber(n int) bool {
	sol1 := -0.5 + math.Sqrt(0.25+float64(2*n))
	sol2 := -0.5 - math.Sqrt(0.25+float64(2*n))
	if sol1 == float64(int(sol1)) || sol2 == float64(int(sol2)) {
		return true
	}
	return false
}

func problem42() {
	data, err := os.ReadFile("input/42.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
	}
	words := strings.Split(string(data), ",")

	count := 0
	for _, word := range words {
		val := 0
		for _, char := range word {
			if char == '"' {
				continue
			}
			val += int(char) - 'A' + 1
		}
		if isTriangleNumber(val) {
			count++
		}
	}
	fmt.Println(count)
}
