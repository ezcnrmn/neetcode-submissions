func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	var maxArea int
	for left < right {
		area := (right-left)*min(heights[left], heights[right])
		maxArea = max(maxArea, area)
		if heights[left]>heights[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
