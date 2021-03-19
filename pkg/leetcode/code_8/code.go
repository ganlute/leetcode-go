package code_8

const (
	max = 1 << 31 -1
	min = -(1 << 31)
)
func myAtoi(s string) int {

	v := 0
	a := 1
	// 先trim掉空格
	for i := range s {
		if string(s[i]) == " " {
			continue
		}
		s = s[i:]
		break
	}
	if len(s) == 0 {
		return 0
	}
	if string(s[0]) == "-" {
		a = -1
		s = s[1:]
	}else if string(s[0]) == "+" {
		s = s[1:]
	}
	for i := range s {
		isNum, number := isNumber(string(s[i]))
		if isNum {
			v = v * 10 + number
		}else {
			break
		}
		if v > max {
			break
		}
	}
	v = v * a
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
func isNumber(v string) (bool, int) {
	numberList := map[string]int{
		"1":1,
		"2":2,
		"3":3,
		"4":4,
		"5":5,
		"6":6,
		"7":7,
		"8":8,
		"9":9,
		"0":0,
	}
	for key, value := range numberList {
		if key == v {
			return true, value
		}
	}
	return false, 0
}
