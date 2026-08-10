package problems

import "fmt"

func init() {
	Register(17, problem17)
}

func problem17() {
	letters := map[int]int{}
	letters[1] = 3   // one
	letters[2] = 3   // two
	letters[3] = 5   // three
	letters[4] = 4   // four
	letters[5] = 4   // five
	letters[6] = 3   // six
	letters[7] = 5   // seven
	letters[8] = 5   // eight
	letters[9] = 4   // nine
	letters[10] = 3  // ten
	letters[11] = 6  // eleven
	letters[12] = 6  // twelve
	letters[13] = 8  //thirteen
	letters[14] = 8  // fourteen
	letters[15] = 7  // fifteen
	letters[16] = 7  // sixteen
	letters[17] = 9  // seventeen
	letters[18] = 8  // eighteen
	letters[19] = 8  // nineteen
	letters[20] = 6  // twenty
	letters[30] = 6  // thirty
	letters[40] = 5  // forty
	letters[50] = 5  // fifty
	letters[60] = 5  // sixty
	letters[70] = 7  // seventy
	letters[80] = 6  // eighty
	letters[90] = 6  // ninety
	letters[100] = 7 // hundred
	and := 3

	onesLetters := 190 * (letters[1] + letters[2] + letters[3] + letters[4] + letters[5] + letters[6] + letters[7] + letters[8] + letters[9])
	teens := 10 * (letters[10] + letters[11] + letters[12] + letters[13] + letters[14] + letters[15] + letters[16] + letters[17] + letters[18] + letters[19])
	tens := 100 * (letters[20] + letters[30] + letters[40] + letters[50] + letters[60] + letters[70] + letters[80] + letters[90])
	ands := 99 * 9 * (and)
	hundred := 900 * letters[100]
	thousand := 3 + 8 // one thousand
	fmt.Println(onesLetters + teens + tens + ands + hundred + thousand)
}
