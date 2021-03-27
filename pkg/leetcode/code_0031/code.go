package code_0031

import "sort"

// 如果nums已经是逆序了，就直接顺序排列
// 否则，找到非逆序的最小区间，首位调大，后面顺序
func nextPermutation(nums []int) {
	if isSorted(nums) {
		sort.Ints(nums)
		return
	}
	i := len(nums) - 1
	for i > 0 {
		if isSorted(nums[i:]) {
			i--
			continue
		} else {
			break
		}
	}
	// 找到了i
	// 从nums[i:]中找到比nums[i]大且最接近的数
	delta := -1
	tmpj := -1
	for j := i + 1; j < len(nums); j++ {
		if tmpj == -1 {
			delta = nums[j] - nums[i]
			tmpj = j
			continue
		}
		if (nums[j] > nums[i]) && ((nums[j] - nums[i]) < delta) {
			delta = nums[j] - nums[i]
			tmpj = j
		}
	}
	nums[i], nums[tmpj] = nums[tmpj], nums[i]
	// 剩下的部分顺序排列
	tNums := nums[i+1:]
	sort.Ints(tNums)
	for j := i + 1; j < len(nums); j++ {
		nums[j] = tNums[j-i-1]
	}
	return
}

func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}
