package code_0056

// 先排序, 按左排序
// 再合并
func merge(intervals [][]int) [][]int {
	// 冒泡排序
	for i := 0; i < len(intervals); i++ {
		found := false
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
				found = true
			}
		}
		if !found {
			break
		}
	}
	// 合并
	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 {
			result = append(result, intervals[i])
		} else {
			tResult := mergeResult(result[len(result)-1], intervals[i])
			result = result[:len(result)-1]
			result = append(result, tResult...)
		}
	}
	return result
}
func mergeResult(intervals1 []int, intervals2 []int) [][]int {
	if intervals2[0] <= intervals1[1] {
		return [][]int{[]int{intervals1[0], max(intervals2[1], intervals1[1])}}
	}
	return [][]int{intervals1, intervals2}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
