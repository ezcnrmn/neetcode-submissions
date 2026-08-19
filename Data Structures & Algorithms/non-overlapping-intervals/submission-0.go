import "slices"

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		prev, cur := result[len(result)-1], intervals[i]
		if prev[1] <= cur[0] {
			result = append(result, cur)
			continue
		}

		if prev[1] - prev[0] > cur[1] - cur[0] {
			result[len(result)-1] = cur
		}
	}
	return len(intervals) - len(result)
}
