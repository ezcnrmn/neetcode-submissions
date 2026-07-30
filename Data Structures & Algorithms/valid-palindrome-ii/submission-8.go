func validPalindrome(s string) bool {
	return f(s, 0)
}

func f(s string, depth int) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			if depth > 0 {
				return false
			}
			if f(s[i+1:j+1], 1) || f(s[i:j], 1) {
				return true
			} else {
				return false
			}
		}
		i++
		j--
	}
	return true
}