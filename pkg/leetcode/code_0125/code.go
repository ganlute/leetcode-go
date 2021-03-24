package code_0125

import "strings"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if !isWordChar(s[i]) {
			i++
			continue
		}
		if !isWordChar(s[j]) {
			j--
			continue
		}
		if strings.ToLower(string(s[i])) == strings.ToLower(string(s[j])) {
			i++
			j--
			continue
		} else {
			return false
		}
	}
	return true

}

func isWordChar(r byte) bool {
	return 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || '0' <= r && r <= '9'
}
