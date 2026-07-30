func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (r+l)/2
		col, row := mid%n, mid/n
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return false
}
