func uniquePaths(row int, col int) int {
	grid := make([][]int, row)
	for i := range grid {
		grid[i] = make([]int, col)
	}

	for i := range col {
		grid[0][i] = 1
	}
	for i := range row {
		grid[i][0] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[row-1][col-1]
}
