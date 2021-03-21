package code_42

// 每根柱子的储水量 = min(柱子左边的最大高度, 柱子右边的最大高度) - 柱子的高度
// maxLeft[i] = maxLeft[i-1]
// maxRight[i] = maxRight[i+1]
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxLeft := make([]int, len(height), len(height))
	maxRight := make([]int, len(height), len(height))
	maxLeft[0] = 0
	maxRight[len(height)-1] = 0
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	result := 0
	for i := range height {
		water := max(0, min(maxLeft[i], maxRight[i])-height[i])
		result = result + water
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
