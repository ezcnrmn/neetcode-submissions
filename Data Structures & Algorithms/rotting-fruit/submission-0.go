func orangesRotting(grid [][]int) int {
	var queue [][3]int
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 2 {
				queue = append(queue, [3]int{row, col, 0})
			}
		}
	}

	minutes := -1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range [...][2]int{{1,0}, {0, 1}, {-1, 0}, {0, -1}} {
			newRow, newCol := cur[0] + d[0], cur[1] + d[1]
			if newRow >= 0 && newRow < len(grid) &&
				newCol >= 0 && newCol < len(grid[0]) &&
				grid[newRow][newCol] == 1 {
				grid[newRow][newCol] = 2
				queue = append(queue, [3]int{newRow, newCol, cur[2]+1})
				minutes = cur[2]+1
			}
		}
	}

	return minutes
}