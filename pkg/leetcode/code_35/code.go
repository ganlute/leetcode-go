package code_35

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) {
		if nums[i] < target {
			i++
		} else {
			break
		}
	}
	return i
}
