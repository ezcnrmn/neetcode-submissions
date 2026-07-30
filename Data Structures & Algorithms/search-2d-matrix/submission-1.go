func searchMatrix(matrix [][]int, target int) bool {
	cols, left, right := len(matrix[0]), 0, len(matrix)*len(matrix[0])
	for left <= right {
		mid := (left + right)/2
		row, col := mid/cols, mid%cols
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return false
}
