package code_0200

// 遍历二维数组
func numIslands(grid [][]byte) int {
	islandNum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 发现新大陆
			if grid[i][j] == '1' {
				islandNum = islandNum + 1
			}
			// 感染新大陆
			infect(grid, i, j)
		}
	}
	return islandNum
}
func infect(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '2'
		if j-1 >= 0 && grid[i][j-1] == '1' {
			infect(grid, i, j-1)
		}
		if i-1 >= 0 && grid[i-1][j] == '1' {
			infect(grid, i-1, j)
		}
		if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
			infect(grid, i, j+1)
		}
		if i+1 < len(grid) && grid[i+1][j] == '1' {
			infect(grid, i+1, j)
		}
	}
	return
}
