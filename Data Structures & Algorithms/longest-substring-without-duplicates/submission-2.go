func lengthOfLongestSubstring(str string) int {
	var left, result int
	window := make(map[byte]struct{})
	for right := 0; right < len(str); right++{
		if _, ok := window[str[right]]; !ok {
			window[str[right]] = struct{}{}
			result = max(result, right-left+1)
			continue
		}

		for left < right {
			if str[left] != str[right] {
				delete(window, str[left])
				left++
			} else{ 
				left++
				break
			}
		}
	}
	return result
}
