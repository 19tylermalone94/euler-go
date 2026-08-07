package problems

import "fmt"

func init() {
	Register(16, problem16)
}

func problem16() {
	n := []int{1}
	for range 1000 {
		carry := 0
		for i := len(n) - 1; i >= 0; i-- {
			prod := n[i]*2 + carry
			carry = prod / 10
			n[i] = prod % 10
		}
		if carry > 0 {
			n = append([]int{carry}, n...)
		}
	}
	sum := 0
	for _, val := range n {
		sum += val
	}
	fmt.Println(sum)
}
