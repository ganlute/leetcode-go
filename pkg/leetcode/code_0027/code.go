package code_0027

func removeElement(nums []int, val int) int {
	index := -1
	for i := range nums {
		if nums[i] == val {
			continue
		}else {
			index = index + 1
			nums[index] = nums[i]
		}
	}
	nums = nums[:index+1]
	return index+1
}