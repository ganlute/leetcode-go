package code_0048

// 找规律

// n = 4
// (3, 1) --> (1, 0)
// (1, 0) --> (0, 2)       (i, k) -- > (k, n-i-1)
// (0, 2) --> (2, 3)
// (2, 3) --> (3, 1)

// 5 -- 2
// 4 -- 2

// 原地旋转矩阵
func rotate(matrix [][]int) {
	n := len(matrix)
	m := n / 2
	// 这个边界
	//（index从0开始）
	// 长度为5，下边界 = 1，右边界 = 2
	// 长度为4，下边界 = 1，有边界 = 1
	for j := 0; j < m; j++ {
		for i := 0; i <= m-1+n%2; i++ {
			// 保留第一象限
			temp := matrix[i][j]
			// 第四象限给第一象限
			matrix[i][j] = matrix[n-j-1][i]
			// 第三象限给第四象限
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			// 第二象限给第三象限
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			// 第一象限给第二象限
			matrix[j][n-i-1] = temp
		}
	}
}
