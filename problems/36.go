package problems

import (
	"euler-go/util"
	"fmt"
	"strconv"
)

func init() {
	Register(36, problem36)
}

func problem36() {
	sum := 0
	for i := range 1000000 {
		if util.IsPalindrome(strconv.Itoa(i)) && util.IsPalindrome(util.ToBinaryString(i)) {
			sum += i
		}
	}
	fmt.Println(sum)
}
