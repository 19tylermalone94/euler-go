package problems

func lcm(nums []int) int {
	// remove/ignore numbers that are factors of others
	for _, val := range nums {
		skip := false
		for _, val2 := range nums {
			if val2%val == 0 {
				skip = true
			}
		}
		if skip {
			continue
		}

	}
	// return product of all distinct prime factors of each number
	return 0
}

func problem5() {

}
