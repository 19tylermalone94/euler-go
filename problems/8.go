package problems

import (
	"fmt"
	"os"
)

func init() {
	Register(8, problem8)
}

func problem8() {
	data, err := os.ReadFile("input/8.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
	}
	digits := []rune(string(data))
	for i := 0; i < 13; i++ {
		fmt.Println(digits[i] - '0')
	}
}
