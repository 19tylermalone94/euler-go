package main

import (
	"euler-go/problems"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]
	fmt.Println(args)
	num, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println("conversion error", err)
		return
	}
	problems.RunProblem(num)
}
