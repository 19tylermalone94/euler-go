package util

func IsPalindrome(n string) bool {
	if len(n) <= 1 {
		return true
	}
	if n[0] != n[len(n)-1] {
		return false
	}
	return IsPalindrome(n[1 : len(n)-1])
}
