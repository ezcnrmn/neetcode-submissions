func characterReplacement(s string, k int) int {
	var left, result int
	var counter [26]int
	for right := 0; right < len(s); right++ {
		counter[s[right]-'A']++
		isValid := validateWindow(counter, k)
		if isValid {
			result = max(result, right-left+1)
			continue
		}

		for !isValid {
			counter[s[left]-'A']--
			left++
			isValid = validateWindow(counter, k)
		}
	}
	return result
}	

func validateWindow(counter [26]int, k int) bool {
	var total, maxAmount int
	for _, c := range counter {
		total += c
		maxAmount = max(maxAmount, c)
	}
	return total - maxAmount <= k
}