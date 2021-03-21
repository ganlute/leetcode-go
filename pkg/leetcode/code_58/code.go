package code_58

func lengthOfLastWord(s string) int {
	// trim space in the end
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			s = s[:i]
		} else {
			break
		}
	}

	// start
	start := 0
	end := len(s) - 1

	for i := range s {
		if string(s[i]) == " " {
			start = i + 1
			end = i + 1
		} else {
			end = i
		}
	}
	return end - start + 1
}
