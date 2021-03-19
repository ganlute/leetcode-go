package code_1

func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if j <= i {
				continue
			}
			if num1 + num2 == target  {
				return []int{i, j}
			}
		}
	}
	return nil
}