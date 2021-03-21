package code_0004

// 先合并数组
// 再求中位数
// 这种解法有点初级
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := make([]int, len(nums1)+len(nums2))
	n := 0
	i := 0
	j := 0
	for {
		if i == len(nums1) {
			break
		}
		if j == len(nums2) {
			break
		}
		if nums1[i] < nums2[j] {
			nums3[n] = nums1[i]
			i++
			n++
		} else {
			nums3[n] = nums2[j]
			j++
			n++
		}

	}
	for i < len(nums1) {
		nums3[n] = nums1[i]
		i++
		n++
	}
	for j < len(nums2) {
		nums3[n] = nums2[j]
		j++
		n++
	}
	n = len(nums3)
	if n%2 == 0 {
		v := nums3[n/2] + nums3[n/2-1]
		return float64(v) / float64(2)
	}
	return float64(nums3[n/2])
}
