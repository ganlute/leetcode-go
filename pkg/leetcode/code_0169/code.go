package code_0169

// 先排序 --> 快排
// https://www.sohu.com/a/246785807_684445
// 1、挖坑法
// 2、指针交换法
// 再查询
func majorityElement(nums []int) int {
	quickSort(nums)
	goal := 0
	if len(nums)%2 == 0 {
		goal = len(nums) / 2
	} else {
		goal = len(nums)/2 + 1
	}
	temp := nums[0]
	count := 0
	for i := range nums {
		if nums[i] == temp {
			count++
			if count >= goal {
				return temp
			}
		} else {
			temp = nums[i]
			count = 1
		}
	}
	return -1
}
// 挖坑法
func quickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	privot := nums[len(nums)/2]
	left := 0
	right := len(nums) - 1

	for left < right {
		for nums[right] >= privot && right > left {
			right--
		}
		if right > left {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		for nums[left] <= privot && right > left {
			left++
		}
		if right > left {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	nums[left] = privot

	quickSort(nums[:left])
	quickSort(nums[left+1:])
}
