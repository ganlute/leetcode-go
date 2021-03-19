package code_15

import "sort"

// Given an array nums of n integers, are there elements a, b,
// c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Notice that the solution set must not contain duplicate triplets.
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	nlen := len(nums)
	if nlen < 3 {
		return result
	}
	// 排序
	sort.Ints(nums)

	for i := 0; i < nlen; i++ {
		start := i + 1
		end := nlen - 1
		for end > start {
			if nums[i]+nums[start]+nums[end] == 0 {
				arr := generateThreeArr(nums[i], nums[start], nums[end])
				if !inResult(arr, result) {
					result = append(result, arr)
				}
				start = start + 1
				end = end - 1
				continue
			}
			if nums[i]+nums[start]+nums[end] > 0 {
				end = end - 1
				continue
			}
			if nums[i]+nums[start]+nums[end] < 0 {
				start = start + 1
				continue
			}
		}
	}
	// 去重
	return result
}
func generateThreeArr(num1, num2, num3 int) []int {
	a := max(max(num1, num2), num3)
	c := min(min(num1, num2), num3)
	b := num1 + num2 + num3 - a - c
	return []int{a, b, c}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func inResult(arr []int, result [][]int) bool {
	for i := range result {
		if result[i][0] == arr[0] &&
			result[i][1] == arr[1] &&
			result[i][2] == arr[2] {
			return true
		}
	}
	return false
}
