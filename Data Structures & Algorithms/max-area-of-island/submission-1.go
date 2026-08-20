func maxAreaOfIsland(grid [][]int) int {
	var result int
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == 1 {
				result = max(countArea(grid, row, col), result)
			}
		}
	}
	return result
}

func countArea(grid [][]int, row, col int) int {
	var area int
	stack := [][2]int{{row, col}}
	neighbours := [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for len(stack) > 0 {
		row, col := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		grid[row][col] = 0
		area++
		for _, n := range neighbours {
			newRow, newCol := row + n[0], col + n[1]
			if newRow >= 0 && newRow < len(grid) &&
				newCol >= 0 && newCol < len(grid[0]) &&
				grid[newRow][newCol] == 1 {

				stack = append(stack, [2]int{newRow, newCol})
			}
		}
	}
	return area
}