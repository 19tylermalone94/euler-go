package problems

import "fmt"

func sumOfSquares(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val * val
	}
	return sum
}

func squareOfSum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum * sum
}

func problem6() {
	nums := make([]int, 100)
	for i := range 100 {
		nums[i] = i + 1
	}
	fmt.Println(squareOfSum(nums) - sumOfSquares(nums))
}
