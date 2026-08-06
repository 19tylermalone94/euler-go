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
	maxProd := 0
	prod := 1
	count := 0
	i := 0
	for {
		if i >= len(digits) {
			break
		}
		if digits[i]-'0' == 0 {
			prod = 1
			count = 0
		} else {
			prod *= int(digits[i] - '0')
			count++
		}

		if count == 14 {
			fmt.Println(digits[i-13] - '0')
			prod /= int(digits[i-13] - '0')
			count--
		}
		if count == 13 {
			maxProd = max(maxProd, prod)
		}
		i++
	}
	fmt.Println(maxProd)
}
