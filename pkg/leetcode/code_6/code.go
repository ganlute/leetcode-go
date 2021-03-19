package code_6

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	groupSize := numRows + numRows - 2
	groupNum := len(s) / groupSize
	if (len(s) % groupSize) != 0 {
		groupNum = groupNum + 1
	}
	result := strings.Builder{}
	for j := 0; j < numRows; j++ {
		for i := 0; i < groupNum; i++ {
			index := i*groupSize + j
			if index >= len(s) {
				continue
			}
			result.WriteByte(s[index])
			if j != 0 && j != numRows-1 {
				index := i*groupSize + groupSize- j
				if index >= len(s) {
					continue
				}
				result.WriteByte(s[index])
			}
		}
	}
	return result.String()
}
