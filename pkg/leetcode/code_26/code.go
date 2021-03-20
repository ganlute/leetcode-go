package code_26

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] != nums[index] {
			index = index + 1
			nums[index] = nums[i]
		}else {
			continue
		}
	}
	nums = nums[:index+1]
	return index + 1
}