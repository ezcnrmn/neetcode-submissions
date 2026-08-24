type StackItem struct {
	Row int
	Col int
	Pac bool
	Ant bool
}

func pacificAtlantic(heights [][]int) [][]int {
	mem := make(map[[2]int][2]bool, len(heights) * len(heights[0]))
	var stack []StackItem
	for col := range len(heights[0])-1 {
		stack = append(stack, StackItem{0, col, true, false})
	}
	for row := 1; row < len(heights)-1; row++ {
		stack = append(stack, StackItem{row, 0, true, false})
	}
	for col := 1; col < len(heights[0]); col++ {
		stack = append(stack, StackItem{len(heights)-1, col, false, true})
	}
	for row := 1; row < len(heights)-1; row++ {
		stack = append(stack, StackItem{row, len(heights[0])-1, false, true})
	}
	stack = append(stack, StackItem{0, len(heights[0])-1, true, true}, StackItem{len(heights)-1, 0, true, true})

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited := mem[[2]int{cur.Row, cur.Col}]
		mem[[2]int{cur.Row, cur.Col}] = [2]bool{visited[0] || cur.Pac, visited[1] || cur.Ant}

		for _, d := range [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			newRow, newCol := cur.Row + d[0], cur.Col + d[1]
			if newRow < 0 || newRow >= len(heights) ||
				newCol < 0 || newCol >= len(heights[0]) ||
				heights[newRow][newCol] < heights[cur.Row][cur.Col] {
				continue
			}

			visited := mem[[2]int{newRow, newCol}]
			if visited[0] && visited[1] || visited[0] == cur.Pac && visited[1] == cur.Ant {
				continue
			}
			stack = append(stack, StackItem{newRow, newCol, cur.Pac, cur.Ant})
		}
	}

	var result [][]int
	for point, flags := range mem {
		if flags[0] && flags[1] {
			result = append(result, point[:])
		}
	}
	return result
}
