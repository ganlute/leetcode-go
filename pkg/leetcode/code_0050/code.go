package code_0050

func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	} else if x == -1 {
		if n%2 == 1 {
			return -1
		} else {
			return 1
		}
	}

	result := float64(1)
	i := 1
	if n < 0 {
		i = -1
		n = -n
	}
	for n > 0 {
		result = result * x
		n--
	}
	if i == -1 {
		result = 1 / result
	}
	return result
}
