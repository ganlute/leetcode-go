package code_0066

func plusOne(digits []int) []int {
	t := 0
	digits[len(digits) - 1] = digits[len(digits) - 1] + 1
	for i := len(digits) - 1; i >= 0; i-- {
		p := (t + digits[i]) % 10
		q := (t + digits[i]) / 10
		digits[i] = p
		t = q
	}
	if t == 0 {
		return digits
	}else {
		result := make([]int, 0)
		result = append(result, t)
		result = append(result, digits...)
		return result
	}
}
