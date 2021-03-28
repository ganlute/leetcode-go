package code_0029

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend < 0 {
		sign = sign * -1
		dividend = - dividend
	}
	if divisor < 0 {
		sign = sign * -1
		divisor = - divisor
	}
	if divisor == 1 {
		result := sign * dividend
		if result < -1<<31 || result > 1<<31-1 {
			return 1<<31 - 1
		} else {
			return result
		}
	}
	i := 0
	for dividend >= divisor {
		i++
		dividend = dividend - divisor
	}
	result := i * sign
	if result < -1<<31 || result > 1<<31-1 {
		return 1<<31 - 1
	} else {
		return result
	}
}
