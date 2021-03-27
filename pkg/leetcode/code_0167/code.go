package code_0167

func twoSum(numbers []int, target int) []int {
	a := 0
	b := len(numbers) - 1
	for i := 0; i < len(numbers)-1; i++ {
		start := i + 1
		end := len(numbers) - 1
		for end >= start {
			mid := (end + start) / 2
			if numbers[i]+numbers[mid] == target {
				a = i
				b = mid
				return []int{a + 1, b + 1}
			} else if numbers[i]+numbers[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	// 一定有解,走不过来
	return []int{a + 1, b + 1}
}
