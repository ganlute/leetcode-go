package code_14

import "strings"

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	result := strings.Builder{}
	if len(strs) == 0 {
		return result.String()
	}
	for i:=0; i < len(strs[0]);i++ {
		b := strs[0][i]
		match := true
		for j := 1; j < len(strs); j ++ {
			if len(strs[j]) <= i {
				match = false
				break
			}
			if b != strs[j][i] {
				match = false
				break
			}
		}
		if match {
			result.WriteByte(b)
		}else {
			break
		}
	}
	return result.String()
}