func lengthOfLongestSubstring(str string) int {
	var left, result int
	window := make(map[byte]struct{})
	for right := 0; right < len(str); right++{
		if _, ok := window[str[right]]; !ok {
			window[str[right]] = struct{}{}
			result = max(result, right-left+1)
		} else {
			for left < right {
				delete(window, str[left])
				left++
				if _, ok := window[str[right]]; !ok {
					break
				}
			}
			window[str[right]] = struct{}{}
		}
	}
	return result
}
