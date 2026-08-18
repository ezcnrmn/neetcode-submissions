func insert(intervals [][]int, newInterval []int) [][]int {
    insertion := newInterval
    result := make([][]int, 0, len(intervals)+1)
    for i, interval := range intervals {
        if interval[1] < insertion[0] {
            result = append(result, interval)
        } else if insertion[1] < interval[0] {
            result = append(result, insertion)
            result = append(result, intervals[i:]...)
            return result
        } else {
            insertion = []int{min(interval[0], insertion[0]), max(interval[1], insertion[1])}
        }
    }
    result = append(result, insertion)
    return result
}