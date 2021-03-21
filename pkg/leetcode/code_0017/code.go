package code_0017

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		switch digits {
		case "2":
			return []string{"a", "b", "c"}
		case "3":
			return []string{"d", "e", "f"}
		case "4":
			return []string{"g", "h", "i"}
		case "5":
			return []string{"j", "k", "l"}
		case "6":
			return []string{"m", "n", "o"}
		case "7":
			return []string{"p", "q", "r", "s"}
		case "8":
			return []string{"t", "u", "v"}
		case "9":
			return []string{"w", "x", "y", "z"}
		}
	}
	left := letterCombinations(string(digits[0]))
	right := letterCombinations(digits[1:])
	result := make([]string, 0)
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			result = append(result, left[i]+right[j])
		}
	}
	return result
}
