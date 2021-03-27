package code_0078

// nums[0] nums[1:]
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}, []int{}}
	}
	left := [][]int{[]int{nums[0]}, []int{}}
	right := subsets(nums[1:])
	result := make([][]int, 0)
	for i := range left {
		for j := range right {
			r := make([]int, 0)
			r = append(r, left[i]...)
			r = append(r, right[j]...)
			result = append(result, r)
		}
	}
	return result
}
