func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0, len(intervals)+1)

	var i int
	for intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	start := min(intervals[i][0], newInterval[0])

	for intervals[i][1] < newInterval[1]  {
		i++
	}

	if intervals[i][0] > newInterval[1] {
		result = append(result, []int{start, newInterval[1]})
	} else {
		result = append(result, []int{start, max(intervals[i][1], newInterval[1])})
		i++
	}

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
