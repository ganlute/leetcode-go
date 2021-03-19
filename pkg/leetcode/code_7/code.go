package code_7

const (
	max = 1 << 31 -1
	min = -(1 << 31)
)

func reverse(x int) int {
	i := 1
	if x < 0 {
		x = -x
		i = -1
	}
	v := int64(0)
	for x > 0 {
		v = v * 10 + int64(x) % 10
		x = x / 10
	}
	v = int64(i) * v
	if v < int64(min) || v > int64(max) {
		return 0
	}

	return int(v)
}