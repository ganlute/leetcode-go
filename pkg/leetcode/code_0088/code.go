package code_0088

// 从小到大会有频繁的数组移动
// 那就从大到小咯
func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m - 1
	q := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if p < 0 || q < 0 {
			break
		}
		if nums2[q] > nums1[p] {
			nums1[i] = nums2[q]
			q--
		} else {
			nums1[i] = nums1[p]
			p--
		}
	}
	if q < 0 {
		return
	}
	for q >= 0 {
		nums1[q] = nums2[q]
		q--
	}
	return
}
