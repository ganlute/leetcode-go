package code_0003

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	maxLen := 0
	for i := 0; i < length; i++ {
		tempContainer := make(map[uint8]int, 0)
		for j := i; j < length; j++ {
			if _, ok := tempContainer[s[j]]; ok {
				break
			} else {
				tempContainer[s[j]] = 1
			}
		}
		maxLen = max(maxLen, len(tempContainer))
	}
	return maxLen
}

func max(v1 int, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
