package code_0046

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	result := make([][]int, 0)
	for i := range nums {
		newNums := make([]int, 0)
		newNums = append(newNums, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		tempArr := permute(newNums)
		for j := range tempArr {
			t := make([]int, 0)
			t = append(t, nums[i])
			t = append(t, tempArr[j]...)
			tempArr[j] = t
		}
		result = append(result, tempArr...)
	}
	return result
}
