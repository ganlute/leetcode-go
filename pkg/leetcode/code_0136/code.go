package code_0136

import "sort"

func singleNumber(nums []int) int {

	sort.Ints(nums)
	for i := 1; i < len(nums); i = i + 2 {
		if nums[i] == nums[i-1] {
			continue
		} else {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}
