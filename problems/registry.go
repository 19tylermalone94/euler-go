package problems

import "fmt"

var probs = map[int]func(){}

func Register(n int, prob func()) {
	probs[n] = prob
}

func RunProblem(i int) {
	prob, ok := probs[i]
	if !ok {
		fmt.Println("Problem not found")
		return
	}
	prob()
}
