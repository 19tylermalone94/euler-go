package util

import (
	"math/big"
	"strconv"
	"strings"
)

func DigitCount(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func BigDigitCount(n *big.Int) int {
	return len(n.String())
}

func ToBinaryString(n int) string {
	b := new(strings.Builder)
	for n > 0 {
		b.WriteString(strconv.Itoa(n % 2))
		n /= 2
	}
	return b.String()
}
