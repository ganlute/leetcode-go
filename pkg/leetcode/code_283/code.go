package code_283

// 保持非零元素的相对顺序
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		found := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				found = true
			}
		}
		if !found {
			break
		}
	}
}
