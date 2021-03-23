package code_0054

import "fmt"

func spiralOrder(matrix [][]int) []int {

	result := make([]int, 0)
	n := len(matrix)    // 行数
	m := len(matrix[0]) // 列数

	for start := 0; start < min(n/2+n%2, m/2+m%2); {
		i := start + 1
		j := start
		// 行固定，列右移  start ---> m - start - 1
		for j < m-start {
			fmt.Println(fmt.Sprintf("%d %d --> %d", start, j, matrix[start][j]))
			result = append(result, matrix[start][j])
			j++
		}
		// 列固定，行下移  start --> n - start - 1
		for i < n-start {
			fmt.Println(fmt.Sprintf("%d %d --> %d", i, m-start-1, matrix[i][m-start-1]))
			result = append(result, matrix[i][m-start-1])
			i++
		}
		i = n - start - 1 - 1
		j = m - start - 1 - 1
		// 行固定，列左移  m - start - 1 -- > start
		for j >= start && (n-start-1 != start) {
			fmt.Println(fmt.Sprintf("%d %d --> %d", n-start-1, j, matrix[n-start-1][j]))
			result = append(result, matrix[n-start-1][j])
			j--
		}
		// 列固定，行上移   n - start - 1 -- > start
		for i > start && (m-start-1 != start) {
			fmt.Println(fmt.Sprintf("%d %d --> %d", i, start, matrix[i][start]))
			result = append(result, matrix[i][start])
			i--
		}
		start++
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
