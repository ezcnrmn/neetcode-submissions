func insert(intervals [][]int, newInterval []int) [][]int {
    all := make([][]int, 0, len(intervals)+1)
    var inserted bool
    for _, interval := range intervals {
        if !inserted && interval[0] > newInterval[0] {
            all = append(all, newInterval, interval)
            inserted = true
            continue
        }
        all = append(all, interval)
    }
    if !inserted {
        all = append(all, newInterval)
    }

    merged := make([][]int, 1, len(all))
    merged[0] = all[0]
    for i := 1; i < len(all); i++ {
        if merged[len(merged)-1][1] >= all[i][0] {
            merged[len(merged)-1][1] =  max(merged[len(merged)-1][1], all[i][1])
        } else {
            merged = append(merged, all[i])
        }
    }

    return merged
}
