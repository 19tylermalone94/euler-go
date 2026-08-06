package problems

import "fmt"

func init() {
	Register(9, problem9)
}

func problem9() {
	for i := 2; i < 1000; i++ {
		for j := 2; j < 1000; j++ {
			for k := 2; k < 1000; k++ {
				if i*i+j*j == k*k && i+j+k == 1000 {
					fmt.Println(i * j * k)
					return
				}
			}
		}
	}
}
