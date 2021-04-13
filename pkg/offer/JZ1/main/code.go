package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */

// 1、暴力解: 遍历数组，match就返回true
// 复杂度：时间复杂度 O(n * n) 空间复杂度 O(1)
func Find( target int ,  array [][]int ) bool {
	for i := 0; i < len(array); i ++ {
		for j := 0; j < len(array[i]); j ++ {
			if array[i][j] == target {
				return true
			}
		}
	}
	return false
}
// 2、二分查找：
// 先查找