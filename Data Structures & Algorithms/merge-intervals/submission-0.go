import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		prev := merged[len(merged)-1]
		if cur[0] <= prev[1] {
			prev[1] = max(cur[1], prev[1])
		} else {
			merged = append(merged, cur)
		}
	}
	return merged
}
