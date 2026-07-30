func characterReplacement(s string, k int) int {
	var left, changed, result int
	for right := 0; right < len(s); right++ {
		if s[left] == s[right] {
			result = max(result, right-left+1)
			continue
		}

		if changed < k {
			changed++
			result = max(result, right-left+1)
			continue
		}

		char := s[left]
		for left < right || changed >= k {
			if s[left] != char {
				left++
				changed--
				break
			}

			left++
		}
	}
	return result
}	
