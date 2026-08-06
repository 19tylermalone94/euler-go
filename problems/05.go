package problems

import (
	"euler-go/utility"
	"fmt"
	"math"
)

func init() {
	Register(5, problem5)
}

func lcm(nums []int) int {
	primePowers := make(map[int]int)
	for _, num := range nums {
		pp := make(map[int]int)
		for _, factor := range utility.PrimeFactors(num) {
			_, ok := pp[factor]
			if ok {
				pp[factor]++
			} else {
				pp[factor] = 1
			}
		}
		for key := range pp {
			_, ok := primePowers[key]
			if ok {
				primePowers[key] = max(primePowers[key], pp[key])
			} else {
				primePowers[key] = pp[key]
			}
		}
	}
	product := 1
	for prime := range primePowers {
		product *= int(math.Pow(float64(prime), float64(primePowers[prime])))
	}
	return product
}

func problem5() {
	n := 20
	nums := make([]int, n)
	for i := range n {
		nums[i] = i + 1
	}
	fmt.Println(lcm(nums))
}
