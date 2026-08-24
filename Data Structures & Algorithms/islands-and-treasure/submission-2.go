func islandsAndTreasure(grid [][]int) {
    var queue [][3]int
    for row := range grid {
        for col := range grid[0] {
            if grid[row][col] == 0 {
                queue = append(queue, [3]int{row, col, 0})
            }
        }
    }

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, d := range [...][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
            newRow, newCol, distance := cur[0] + d[0], cur[1] + d[1], cur[2] + 1
            if newRow < 0 || newRow >= len(grid) ||
                newCol < 0 || newCol >= len(grid[0]) ||
                grid[newRow][newCol] != math.MaxInt32 {
                continue
            }

            grid[newRow][newCol] = distance
            queue = append(queue, [3]int{newRow, newCol, distance})
        }
    }
}
