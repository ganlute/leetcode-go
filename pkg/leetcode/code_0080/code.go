package code_0080

// 先排序
// 再统计
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 冒泡排序
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	// 已经排序
	start := 0
	tmp := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == tmp {
			count++
		} else {
			count = 1
			tmp = nums[i]
		}
		if count <= 2 {
			nums[start] = nums[i]
			start++
		}
	}
	return start
}
