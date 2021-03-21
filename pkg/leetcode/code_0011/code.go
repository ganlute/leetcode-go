package code_0011

// 双指针法
// start = 0
// end = len - 1
// 怎么移动？
// 每次选择高度小的去移动
// 为什么？
// 如果每次选择高度大的去移动，面积一定变小
// 如果每次选择高度小的去移动，面积可能增大

func maxArea(height []int) int {
	maxValue := 0
	start := 0
	end := len(height) - 1
	for end > start {
		maxValue = max(maxValue, min(height[end], height[start])*(end-start))
		if height[end] > height[start] {
			start = start + 1
		} else {
			end = end - 1
		}
	}
	return maxValue
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
