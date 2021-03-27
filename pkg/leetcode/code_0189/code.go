package code_0189

func rotate(nums []int, k int) {
	k = k % len(nums)
	newNums := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := range nums {
		nums[i] = newNums[i]
	}
}
