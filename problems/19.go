package problems

import "fmt"

func init() {
	Register(19, problem19)
}

func problem19() {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	d := 7
	m := 12

	day := 1
	date := 1
	month := 0
	year := 1901

	count := 0

	for {
		if year == 2000 && date == 2 {
			break
		}
		if date == 1 && day == 0 {
			count++
		}
		if month == 1 && year%4 == 0 && date == 29 {
			continue
		}
		day = (day + 1) % d
		date++
		if date > months[month] {
			month++
			date = 1
		}
		if month >= m {
			year++
			month = 0
		}
	}
	fmt.Println(count)
}
