package code_0028

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i:=0; i < len(haystack); i ++ {
		if isPrefix(needle, haystack[i:]) {
			return i
		}
	}
	return -1
}

func isPrefix(prefix, str string) bool {
	i:= 0
	if len(prefix) > len( str) {
		return false
	}
	for {
		if i >= len(prefix) {
			break
		}
		if prefix[i] == str[i] {
			i++
			continue
		}else {
			return false
		}
	}
	return true
}