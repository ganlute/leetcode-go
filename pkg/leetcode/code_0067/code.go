package code_0067

func addBinary(a string, b string) string {
	index_a := len(a) - 1
	index_b := len(b) - 1
	result := ""
	tValue := 0
	for index_a >= 0 && index_b >= 0 {
		tValue = transferStrToInt(a[index_a]) + transferStrToInt(b[index_b]) + tValue
		result = transferIntToStr(tValue%2) + result
		tValue = tValue / 2
		index_a--
		index_b--
	}
	for index_a >= 0 {
		tValue = transferStrToInt(a[index_a]) + tValue
		result = transferIntToStr(tValue%2) + result
		tValue = tValue / 2
		index_a--
	}
	for index_b >= 0 {
		tValue = transferStrToInt(b[index_b]) + tValue
		result = transferIntToStr(tValue%2) + result
		tValue = tValue / 2
		index_b--
	}
	if tValue != 0 {
		result = transferIntToStr(tValue) + result
	}
	return result
}
func transferIntToStr(a int) string {
	if a == 0 {
		return "0"
	} else if a == 1 {
		return "1"
	} else {
		return ""
	}
}

func transferStrToInt(in byte) int {
	if in == '1' {
		return 1
	} else {
		return 0
	}
}
