import "slices"

/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	dots := make([][2]int, 0, len(intervals)*2)
	for _, interval := range intervals {
		dots = append(dots, [2]int{interval.start, 1}, [2]int{interval.end, -1})
	}
	slices.SortFunc(dots, func(a, b [2]int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	var cur, maxRooms int
	for _, dot := range dots {
		cur += dot[1]
		maxRooms = max(maxRooms, cur)
	}
	return maxRooms
}
