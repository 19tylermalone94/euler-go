package util

func IsRightTriangle(a, b, c int) bool {
	if a < 1 || b < 1 || c < 1 {
		return false
	}
	return a*a+b*b == c*c
}
