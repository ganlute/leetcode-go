package code_0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						if len(result) == 0 {
							result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
						} else {
							if isInResult(result, []int{nums[i], nums[j], nums[k], nums[l]}) {
								continue
							} else {
								result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
							}
						}

					}
				}
			}
		}
	}
	return result
}
func isInResult(result [][]int, arr []int) bool {
	for i := range result {
		if result[i][0] == arr[0] &&
			result[i][1] == arr[1] &&
			result[i][2] == arr[2] &&
			result[i][3] == arr[3] {
			return true
		}
	}
	return false
}
