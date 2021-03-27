package code_0069

// 在 0 ～ x 中查找一个数
// 这个数的平方<x
// 这个数+1的平方>x
// 二分查找

func mySqrt(x int) int {
	start := 0
	end := x
	for end >= start {
		mid := (start + end) / 2
		v := mid * mid
		if v < x {
			start = mid + 1
		} else if v == x {
			return mid
		} else {
			end = mid - 1
		}
	}
	return end
}
