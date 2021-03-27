package code_0075

// 冒泡排序
func sortColors(nums []int) {
	m := len(nums)
	for j := m - 1; j >= 0; j-- {
		found := false
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				found = true
			}
		}
		if !found {
			break
		}
	}
	return
}
