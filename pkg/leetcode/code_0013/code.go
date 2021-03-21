package code_0013

func romanToInt(s string) int {
	result := 0
	for i:=0; i< len(s);i++ {
		if i+1 == len(s) {
			result = result + getValue(s[i])
			break
		}
		if getValue(s[i]) >= getValue(s[i+1]) {
			result = result + getValue(s[i])
		}else {
			result = result - getValue(s[i])
		}
	}
	return result
}

func getValue(uint uint8) int {
	switch string(uint) {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}


