package code_0704

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for end >= start {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
