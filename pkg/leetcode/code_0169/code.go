package code_0169

// 先排序 --> 快排
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

func quickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	left := 0
	// 从最左边挖个洞
	privot := nums[left]
	right := len(nums) - 1

	for left < right {
		// 左边有个洞，从右边开始
		for nums[right] > privot && right > left {
			right--
		}
		if right > left {
			// 右边空了个洞
			nums[left] = nums[right]
			left++
		}
		for nums[left] <= privot && right > left {
			left++
		}
		if right > left {
			nums[right] = nums[left]
			right--
		}
	}
	nums[left] = privot
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}
