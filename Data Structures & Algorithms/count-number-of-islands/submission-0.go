func numIslands(grid [][]byte) int {
	var counter int
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == '1' {
				counter++
				paint(grid, row, col)
			}
		}
	}
	return counter
}

func paint(grid [][]byte, row, col int) {
	var queue [][2]int = [][2]int{{row, col}}
	neighbours := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		grid[row][col] = '0'
		for _, n := range neighbours {
			if row + n[0] >= 0 && row + n[0] < len(grid) &&
				col + n[1] >= 0 && col + n[1] < len(grid[0]) &&
				grid[row+n[0]][col+n[1]] == '1' {
				queue = append(queue, [2]int{row+n[0], col+n[1]})
			}
		}
	}
}