package code_0059

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	// 初始化
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	num := 1
	// 几轮循环
	for i := 0; i <= n/2; i++ {

		// 每一轮循环做什么
		for j := i; j < n-i; j++ {
			matrix[i][j] = num
			fmt.Printf("matrix[%d][%d]=%d\n", i, j, num)
			num++
		}
		for ti := i + 1; ti < n-i; ti++ {
			matrix[ti][n-i-1] = num
			fmt.Printf("matrix[%d][%d]=%d\n", ti, n-i-1, num)
			num++
		}
		for tj := n - i - 2; tj >= i; tj-- {
			matrix[n-i-1][tj] = num
			fmt.Printf("matrix[%d][%d]=%d\n", n-i-1, tj, num)
			num++
		}
		for ti := n - i - 2; ti > i; ti-- {
			matrix[ti][i] = num
			fmt.Printf("matrix[%d][%d]=%d\n", ti, i, num)
			num++
		}
	}
	return matrix
}
