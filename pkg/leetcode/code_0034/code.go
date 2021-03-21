package code_0034

func searchRange(nums []int, target int) []int {
	min:= -1
	max:= -1
	for i := range nums{
		if nums[i] != target {
			continue
		}
		if min == -1 {
			min = i
		}
		max = i
	}
	return []int{min,max}
}