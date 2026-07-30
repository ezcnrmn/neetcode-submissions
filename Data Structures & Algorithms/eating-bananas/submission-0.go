func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	result := right

	for left <= right {
		mid := left + (right-left)/2
		var time int
		for _, pile := range piles {
			if pile % mid == 0 {
				time += pile/mid
			} else {
				time += pile/mid+1
			}
		}

		if time <= h {
			result = mid
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return result
}