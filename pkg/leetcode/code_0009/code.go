package code_0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tempx := x
	y := 0
	for x > 0 {
		y = 10 * y + (x % 10 )
		x = x / 10
	}
	if y == tempx {
		return true
	}
	return false
}