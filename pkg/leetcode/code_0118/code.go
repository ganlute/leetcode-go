package code_0118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	resultArr := make([][]int, 0)
	resultArr = append(resultArr, []int{1})
	for i := 1; i < numRows; i++ {
		resultList := make([]int, i+1)
		resultList[0] = 1
		resultList[i] = 1
		lastList := resultArr[len(resultArr)-1]
		for j := 0; j < len(lastList)-1; j++ {
			resultList[j+1] = lastList[j] + lastList[j+1]
		}
		resultArr = append(resultArr, resultList)
	}
	return resultArr
}
