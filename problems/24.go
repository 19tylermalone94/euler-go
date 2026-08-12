package problems

import (
	"fmt"
	"os"
)

func init() {
	Register(24, problem24)
}

var n int = 10
var count int = 0

func permute(nums []int, perm []int) {
	if len(perm) == n {
		count++
		if count == 1000000 {
			fmt.Println(perm)
			os.Exit(0)
		}
		return
	}
	for i := range nums {
		perm = append(perm, nums[i])
		nums = append(nums[0:i], nums[i+1:]...)
		permute(nums, perm)
		nums = append(nums[0:i], append([]int{perm[len(perm)-1]}, nums[i:]...)...)
		perm = perm[0 : len(perm)-1]
	}
}

func problem24() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permute(nums, []int{})
}
