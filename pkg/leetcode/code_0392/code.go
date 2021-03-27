package code_0392

func isSubsequence(s string, t string) bool {
	j := 0
	for i := range t {
		if j == len(s) {
			break
		}
		if t[i] == s[j] {
			j++
		} else {
			// do nothing
		}
	}
	if j == len(s) {
		return true
	}
	return false
}
