func plusOne(digits []int) []int {
	flag := true
	for i := len(digits)-1; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] -= 10
		} else {
			flag = false
			break
		}
	}
	if !flag {
		return digits
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	copy(result[1:], digits)
	return result
}

