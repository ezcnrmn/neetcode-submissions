func isValidSudoku(board [][]byte) bool {
	var counter map[byte]struct{}

	for row := 0; row < 9; row++ {
		counter = make(map[byte]struct{}, 9)
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			if _, ok := counter[board[row][col]]; ok {
				return false
			}
			counter[board[row][col]] = struct{}{}
		}
	}

	for col := 0; col < 9; col++ {
		counter = make(map[byte]struct{}, 9)
		for row := 0; row < 9; row++ {
			if board[row][col] == '.' {
				continue
			}
			if _, ok := counter[board[row][col]]; ok {
				return false
			}
			counter[board[row][col]] = struct{}{}
		}
	}

	boxes := [9]struct{
		x int
		y int
	}{{1,1}, {1,4}, {1,7}, {4,1}, {4,4}, {4,7}, {7,1}, {7,4}, {7,7}}
	difs := [9]struct{
		x int
		y int
	}{{0,0}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for _, box := range boxes {
		counter = make(map[byte]struct{}, 9)
		for _, dif := range difs {
			x, y := box.x+dif.x, box.y+dif.y
			if board[x][y] == '.' {
				continue
			}
			if _, ok := counter[board[x][y]]; ok {
				return false
			}
			counter[board[x][y]] = struct{}{}
		}
	}

	return true
}
