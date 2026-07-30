func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	err := false
	for i < j {
		if s[i] != s[j] {
			if err {
				return false
			}
			if s[i+1] == s[j] {
				i++
			} else if s[i] == s[j-1] {
				j--
			} else {
				return false
			}
			err = true
		}
		i++
		j--
	}
	return true
}