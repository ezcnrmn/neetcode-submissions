/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import "slices"

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.start - b.start
	})
	for i := 1; i < len(intervals); i++ {
		cur, prev := intervals[i], intervals[i-1]
		if cur.start <= prev.end {
			return false
		}
	}
	return true
}
