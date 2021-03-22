package code_0016

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[2]
	minSize := abs(result, target)
	// 先排序
	sort.Ints(nums)
	// 再遍历
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(target, sum) < minSize {
					result = sum
					minSize = abs(target, sum)
				}
			}
		}
	}
	return result
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
