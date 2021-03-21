package code_38

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	countResult := countAndSay(n - 1)
	// say count
	tempNum := "-1"
	tempCount := 0
	result := ""
	for i := 0; i < len(countResult); i++ {
		if tempNum == string(countResult[i]) {
			tempCount++
			continue
		}
		if tempNum != string(countResult[i]) {
			if tempCount != 0 {
				result = result + fmt.Sprintf("%d%s", tempCount, tempNum)
			}
			tempNum = string(countResult[i])
			tempCount = 1
		}
	}
	if tempCount != 0 {
		result = result + fmt.Sprintf("%d%s", tempCount, tempNum)
	}
	return result
}
