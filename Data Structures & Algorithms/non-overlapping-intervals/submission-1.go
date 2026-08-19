import "slices"

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	lastEnd := intervals[0][1]
	var result int
	for i := 1; i < len(intervals); i++ {
		if lastEnd <= intervals[i][0] {
			lastEnd = intervals[i][1]
		} else {
			result++
		}
	}
	return result
}
