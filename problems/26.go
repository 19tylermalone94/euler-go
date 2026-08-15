package problems

import "fmt"

func init() {
	Register(26, problem26)
}

func period(d int) int {
	m := d
	for m%2 == 0 {
		m /= 2
	}
	for m%5 == 0 {
		m /= 5
	}
	if m == 1 {
		return 0
	}
	r, k := 1, 0
	for {
		r = r * 10 % m
		k++
		if r == 1 {
			return k
		}
	}
}

func problem26() {
	maxPeriod := 0
	maxD := 0
	d := 2
	for {
		if d == 1000 {
			break
		}
		p := period(d)
		if p > maxPeriod {
			maxPeriod = p
			maxD = d
		}
		d++
	}
	fmt.Println(maxD, maxPeriod)
}
